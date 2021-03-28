package auth

import(
	"net/http"
	"github.com/seanaye/geog483-final/server/pkg/redis"
	"github.com/seanaye/geog483-final/server/pkg/jwt"
)

func Middleware(db *redis.RedisService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Header.Get("Authorization")
			

			// allow unauthenticated users
			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			token, err := jwt.ValidateToken(c)

			if err != nil {
				http.Error(w, "Invalid authentication token", http.StatusForbidden)
				return
			}


		}
	}
}
