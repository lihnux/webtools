package server

import (
	"net/http"
	"path"
)

type Routers struct {
	FileServ http.Handler
}

func NewRouters() *Routers {
	routers := &Routers{}
	routers.FileServ = http.FileServer(http.Dir("static"))

	return routers
}

func isStaticFile(r *http.Request) bool {
	switch path.Ext(r.URL.Path) {
	case ".html", ".js", ".css":
		return true;
	}
	return false;
}
func (routers *Routers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isStaticFile(r) || r.URL.Path == "/" {
		routers.FileServ.ServeHTTP(w, r)
	} else {
		// 普通请求处理
	}
}
