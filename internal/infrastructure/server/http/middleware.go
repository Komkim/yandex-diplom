package router

import (
	"github.com/go-chi/render"
	"net/http"
)

func (t *Router) AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		token, err := r.Cookie("token")
		if err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		if token == nil {
			render.Render(w, r, ErrNotAuthenticated)
			return
		}

		_, ok := t.auth.FetchAuth(token.Value)
		if !ok {
			render.Render(w, r, ErrNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
