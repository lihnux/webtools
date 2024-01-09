package server

import (
	"net/http"
)

var (
	Addr = ":80"
)

func Server() *http.Server {
//	r := &Router{
//		FileServ: http.FileServer(http.Dir("static")),
//	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	
	server := &http.Server{
		Addr: Addr,
//		Handler: r,
	}
	return server
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