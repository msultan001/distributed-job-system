package api

import "github.com/go-chi/chi/v5"

type Router struct {
	handler *JoBHandler
	router  chi.Router
}

func NewRouter(h *JoBHandler) *Router {
	return &Router{
		handler: h,
		router:  chi.NewRouter(),
	}
}

func (rr *Router) RegisterRoutes() {
	rr.router.Post("/jobs", rr.handler.CreateJobPost)
	rr.router.Get("/posts/{id}", rr.handler.GetProcessedJob)
}
