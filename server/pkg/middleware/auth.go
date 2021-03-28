package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/seanaye/geog483-final/server/pkg/jwt"
	"github.com/seanaye/geog483-final/server/pkg/redis"
)
type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware(db *redis.RedisService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := r.Header.Get("Authorization")
			

			// allow unauthenticated users
			if c == "" {
				next.ServeHTTP(w, r)
				return
			}

			// turn bearer string into token
			segments := strings.Split(c, " ")

			if segments[0] != "bearer" {
				http.Error(w, "Invalid token", http.StatusForbidden)
			}

			token, err := jwt.ValidateToken(segments[1])

			if err != nil {
				http.Error(w, "Invalid authentication token", http.StatusForbidden)
				return
			}

			id := token.Claims["id"]
			user, err := db.GetUser(id)
			if err != nil {
				http.Error(w, "Could not find user in database", http.StatusNotFound)
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
