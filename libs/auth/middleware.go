package auth

import (
	"encoding/json"
	"errorwrapper"
	"io"
	"net/http"
	"strings"
)

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

var requestData struct {
	OperationName string `json:"operationName"`
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
