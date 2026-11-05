package api

import (
	"log"
	"net/http"
)

type API struct {
	Address string
	r       *Router
}

func NewAPI(address string, router *Router) *API {
	return &API{
		Address: address,
		r:       router,
	}
}

func (a *API) Run() error {
	log.Println("started running server on port", a.Address)
	server := &http.Server{
		ReadTimeout:  10,
		WriteTimeout: 5,
		Handler:      a.r.router,
		Addr:         a.Address,
	}
	return server.ListenAndServe()
}
