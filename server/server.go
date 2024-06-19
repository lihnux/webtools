package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Server struct {
	routers *Routers
	serv *http.Server
}

func NewServer(addr string) *Server {
	server := &Server{}
	server.routers = NewRouters()
	server.serv = &http.Server{
		Addr: addr,
		Handler: server.routers,
	}

	return server
}

func (s *Server) Run() {
	slog.Info(fmt.Sprintf("Http server listening at %s", s.serv.Addr))
	err := s.serv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		slog.Error("Http server starts failed", "reason", err.Error())
		os.Exit(1)
	}
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.TODO(), 20 * time.Second)
	defer cancel()
	if err := s.serv.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed:", "error", err.Error())
	}
}

/*
package main

import (
"log"
"net/http"
)

func main() {
fs := http.FileServer(http.Dir("/usr/src/webtools/static"))
http.Handle("/", fs)

log.Println("Listening on :80...")
err := http.ListenAndServe(":80", nil)
if err != nil {
log.Fatal(err)
}
}
*/