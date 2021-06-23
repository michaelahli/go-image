package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/michaelahli/Simple-Image-Transfer/server/src/controllers"
)

type router struct {
	ctrl controllers.Controller
}

type Router interface {
	RouterGroup() http.Handler
}

func SetupRouter(ctrl controllers.Controller) Router {
	return &router{ctrl: ctrl}
}

func (path *router) RouterGroup() http.Handler {
	routes := chi.NewRouter()
	routes.Group(func(r chi.Router) {
		r.Post("/api/upload-image", path.ctrl.ImageHandler)
		r.Get("/api/check-metadata", path.ctrl.CheckMetadata)
	})
	return routes
}
