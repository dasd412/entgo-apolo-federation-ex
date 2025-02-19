package auth

import (
	"auth"
	"errorwrapper"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// 인증이 필요 없는 API 목록
var publicOperations = map[string]bool{
	"signup":             true,
	"login":              true,
	"IntrospectionQuery": true, //graphql playground 용
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			operationName, _ := auth.ApiOperationNameFromContext(r.Context())

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
