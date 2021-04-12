package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/seanaye/geog483-final/server/pkg/jwt"
	"github.com/seanaye/geog483-final/server/pkg/redis"
	"github.com/seanaye/geog483-final/server/pkg/user"
)

type JSONErrorMessage struct {
	Message string `json:"message"`
}

func getAndValidateUser(db redis.RedisService, token string) (*user.UserItem, error) {

	// turn bearer string into token
	segments := strings.Split(token, " ")

	if segments[0] != "Bearer" {
		return nil, errors.New("Invalid Token")
	}

	claims, valid := jwt.ValidateClaims(segments[1])

	if !valid {
		return nil, errors.New("Invalid Token")
	}

	id := claims["id"].(string)
	user, err := db.GetUser(id)
	if err != nil {
		return nil, errors.New("Could not find user")
	}

	return user, nil
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

			user, err := getAndValidateUser(db, c)

			if err != nil {
				JSONError(w, JSONErrorMessage{"Invalid authentication token"}, http.StatusForbidden)
			}

			ctx := context.WithValue(r.Context(), "user", user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		})
	}
}

func WSInit(service redis.RedisService) func(context.Context, transport.InitPayload) (context.Context, error) {
	return func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
		if initPayload == nil {
			return ctx, nil
		}

		var gg map[string]interface{}
		gg = initPayload
		headers := gg["headers"]
		if headers == nil {
			return ctx, nil
		}
		token := headers.(map[string]interface{})["Authorization"]
		//header := initPayload["headers"]["Authorization"]
		if token == nil {
			return ctx, nil
		}

		if token == "" {
			return ctx, nil
		}
		user, err := getAndValidateUser(service, token.(string))
		if err != nil {
			return nil, err
		}

		// put it in context
		userCtx := context.WithValue(ctx, "user", user)

		// and return it so the resolvers can see it
		return userCtx, nil
	}
}

func ForContext(ctx context.Context) *user.UserItem {
	raw, _ := ctx.Value("user").(*user.UserItem)
	return raw
}
