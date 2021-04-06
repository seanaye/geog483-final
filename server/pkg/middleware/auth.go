package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/seanaye/geog483-final/server/pkg/jwt"
	"github.com/seanaye/geog483-final/server/pkg/redis"
	"github.com/seanaye/geog483-final/server/pkg/user"
)
type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

type JSONErrorMessage struct {
	Message string `json:"message"`
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(err)
}

func Auth(db redis.RedisService) func(http.Handler) http.Handler {
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
				JSONError(w, JSONErrorMessage{"Invalid token"}, http.StatusForbidden)
				return
			}

			claims, valid := jwt.ValidateClaims(segments[1])

			if !valid {
				JSONError(w, JSONErrorMessage{"Invalid authentication token"}, http.StatusForbidden)
				return
			}

			id := claims["id"].(string)
			user, err := db.GetUser(id)
			if err != nil {
				JSONError(w, JSONErrorMessage{"Could not find user in database"}, http.StatusNotFound)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		})
	}
}

func ForContext(ctx context.Context) *user.UserItem {
	raw, _ := ctx.Value(userCtxKey).(*user.UserItem)
	return raw
}
