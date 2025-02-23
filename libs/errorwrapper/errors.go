package errorwrapper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// WrapError  발생한 에러를 GraphQL 오류 형식에 맞게 wrapping합니다.
func WrapError(ctx context.Context, err error) *gqlerror.Error {
	// 기본 httperror presenter로 변환
	gqlErr := graphql.DefaultErrorPresenter(ctx, err)

	var httpErr *HTTPError

	if errors.As(err, &httpErr) {
		gqlErr.Message = httpErr.Message

		switch httpErr.StatusCode {
		case http.StatusBadRequest:
			gqlErr.Extensions = map[string]interface{}{
				"code": "BAD_REQUEST",
			}
		case http.StatusUnauthorized:
			gqlErr.Extensions = map[string]interface{}{
				"code": "UNAUTHORIZED",
			}
		case http.StatusForbidden:
			gqlErr.Extensions = map[string]interface{}{
				"code": "FORBIDDEN",
			}
		case http.StatusNotFound:
			gqlErr.Extensions = map[string]interface{}{
				"code": "NOT_FOUND",
			}
		default:
			gqlErr.Extensions = map[string]interface{}{
				"code": "INTERNAL_ERROR",
			}
		}
	}

	return gqlErr
}

// SetErrorResponse HTTPError를 받아 GraphQL 에러 포맷의 JSON 응답을 작성합니다.
func SetErrorResponse(w http.ResponseWriter, ctx context.Context, err error) {
	gqlErr := WrapError(ctx, err)

	var status int

	var httpErr *HTTPError

	if errors.As(err, &httpErr) {
		status = httpErr.StatusCode
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"errors": []interface{}{gqlErr},
	})
}
