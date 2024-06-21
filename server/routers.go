package server

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"webtools/model"
)

type Routers struct {
	FileServ http.Handler
	SqliteUI http.Handler
}

func NewRouters() *Routers {
	routers := &Routers{}
	routers.FileServ = http.FileServer(http.Dir("static"))

	sqliteUIBackend, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		slog.Error("Parse sqlite ui url failed", "reason", err.Error())
	} else {
		routers.SqliteUI = httputil.NewSingleHostReverseProxy(sqliteUIBackend)
	}
	http.HandleFunc("/stocks", handleStockApi)

	return routers
}

func isStaticFile(r *http.Request) bool {
	switch path.Ext(r.URL.Path) {
	case ".html", ".js", ".css":
		return true
	}
	return false
}
func (routers *Routers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host == "stock.sqlite.local" {
		routers.SqliteUI.ServeHTTP(w, r)
	} else {
		if isStaticFile(r) || r.URL.Path == "/" {
			routers.FileServ.ServeHTTP(w, r)
		} else {
			// 普通请求处理
			http.DefaultServeMux.ServeHTTP(w, r)
		}
	}
}

func handleStockApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		stocks := model.ListStocks()

		// 设置响应头
		w.Header().Set("Content-Type", "application/json")

		// 将响应数据编码为 JSON
		json.NewEncoder(w).Encode(stocks)
	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Error reading JSON request", "reason", err.Error())
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		var stock model.Stock
		err = json.Unmarshal(body, &stock)
		if err != nil {
			slog.Error("Error decoding JSON request", "reason", err.Error())
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		err = stock.Save()
		if err != nil {
			slog.Error("Error saving new stock", "reason", err.Error())
			http.Error(w, "Error saving new stock", http.StatusBadRequest)
			return
		}
		w.Write([]byte("OK"))

	default:
		http.Error(w, "Do not support this method", http.StatusMethodNotAllowed)
	}

}
