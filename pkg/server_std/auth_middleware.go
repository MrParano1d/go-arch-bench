package serverstd

import (
	"context"
	"net/http"

	"github.com/mrparano1d/archs-go/pkg/session"
)

func AuthMiddleware(sessionManager *session.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			sess, err := sessionManager.Get(r.Context(), r.Header.Get("X-Session-ID"))
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", sess.User)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
