package auth

import (
	"auth"
	"context"
	"encoding/json"
	"errors"
	"errorwrapper"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"strings"
)

var requestData struct {
	OperationName string `json:"operationName"`
}

func ApiOperationNameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		operationName, err := getOperationName(r)

		if err != nil {
			errorwrapper.SetErrorResponse(w, r.Context(), &errorwrapper.HTTPError{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid operation: " + err.Error(),
			})
			return
		}

		// rule.go의 AllowIfSignupOrLogin() 등에서 API 이름에 따라 세밀하게 조정하기 위함.
		ctx := WithApiOperationName(r.Context(), operationName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getOperationName(r *http.Request) (string, error) {
	// graphql 요청 본문 (json) 읽기
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	//json 파싱
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return "", err
	}

	//요청 본문을 다시 복원
	r.Body = io.NopCloser(strings.NewReader(string(body)))

	return requestData.OperationName, nil
}

type apiOperationNameContextKey string

const apiOperationNameKey apiOperationNameContextKey = "apiOperationName"

func WithApiOperationName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, apiOperationNameKey, name)
}

func ApiOperationNameFromContext(ctx context.Context) (string, error) {
	operationName, ok := ctx.Value(apiOperationNameKey).(string)
	if !ok {
		return "", errors.New("OperationName not found in context")
	}

	return operationName, nil
}

// 인증이 필요 없는 API 목록
var publicOperations = map[string]bool{
	"signup":             true,
	"login":              true,
	"IntrospectionQuery": true, //graphql playground 용
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			operationName, _ := ApiOperationNameFromContext(r.Context())

			// 인증이 필요없는 요청이면 JWT 검증을 건너뛰고 API 이름을 ctx에 담음.
			if publicOperations[operationName] {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				errorwrapper.SetErrorResponse(w, r.Context(), &errorwrapper.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Authorization header missing",
				})
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := ValidateJwt(tokenString, false)

			if err != nil {
				errorwrapper.SetErrorResponse(w, r.Context(), &errorwrapper.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token",
				})
				return
			}

			//사용자 id 추출
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {
				errorwrapper.SetErrorResponse(w, r.Context(), &errorwrapper.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token claims",
				})
				return
			}

			userId, ok := claims["sub"].(string)
			if !ok {
				errorwrapper.SetErrorResponse(w, r.Context(), &errorwrapper.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token subject",
				})
				return
			}

			role, ok := claims["role"].(string)

			if !ok {
				errorwrapper.SetErrorResponse(w, r.Context(), &errorwrapper.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token role",
				})
				return
			}

			// 사용자 id 및 권한을 context에 저장
			ctx := auth.WithUserId(r.Context(), userId)
			ctx = auth.WithUserAuthority(ctx, auth.NewAuthority(auth.ConvertRole(role)))
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
