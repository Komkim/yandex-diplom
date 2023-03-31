package router

import (
	"github.com/go-chi/render"
	"net/http"
	"yandex-diplom/internal/infrastructure/server/http/response"
)

func (t *Router) AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		token, err := r.Cookie("token")
		if err != nil {
			render.Render(w, r, response.ErrInternalServer(err))
			return
		}

		if token == nil {
			render.Render(w, r, response.ErrNotAuthenticated)
			return
		}

		_, ok := t.auth.FetchAuth(token.Value)
		if !ok {
			render.Render(w, r, response.ErrNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}