package server

import "net/http"

type Router struct {
	FileServ http.Handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		router.FileServ.ServeHTTP(w, r)
	} 
}