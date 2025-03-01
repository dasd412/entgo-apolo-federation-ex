package auth

import (
	"context"
	"net/http"
)

type userPassportContextKey string

const passportKey userPassportContextKey = "userPassportContextKey"

type Passport struct {
	UserId        string
	UserAuthority UserAuthority
}

func PassportMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("user-id")
		role := r.Header.Get("role")

		ctx := context.WithValue(
			r.Context(),
			passportKey,
			Passport{
				UserId:        userID,
				UserAuthority: NewAuthority(ConvertRole(role)),
			})

		// 다음 핸들러 실행
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PassportFromContext(ctx context.Context) Authority {
	v, _ := ctx.Value(passportKey).(Authority)
	return v
}
