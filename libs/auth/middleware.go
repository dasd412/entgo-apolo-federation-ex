package auth

import (
	"context"
	"net/http"
	"strconv"
)

type userPassportContextKey string

const passportKey userPassportContextKey = "userPassportContextKey"

type Passport struct {
	UserId        int
	UserAuthority UserAuthority
}

func PassportMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIdStr := r.Header.Get("user-id")
		role := r.Header.Get("role")

		if userIdStr == "" || role == "" {
			next.ServeHTTP(w, r)
			return
		}

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, "Invalid user id", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			passportKey,
			Passport{
				UserId:        userId,
				UserAuthority: NewAuthority(ConvertRole(role)),
			})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PassportFromContext(ctx context.Context) Passport {
	v, _ := ctx.Value(passportKey).(Passport)
	return v
}
