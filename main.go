package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webtools/server"
)

func main() {
	serv := server.Server()
	
	go func() {
		slog.Info(fmt.Sprintf("Http server listening at %s", server.Addr))
		err := serv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			slog.Error("Http server starts failed", "reason", err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Warn("Server shuting down ...")

	ctx, cancel := context.WithTimeout(context.TODO(), 20 * time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed:", "error", err.Error())
	}

	slog.Warn("Server exited")
}

