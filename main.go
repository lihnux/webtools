package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"webtools/server"
)

func main() {
	serv := server.NewServer(":443")
	go serv.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Warn("Server shuting down ...")

	serv.Stop()

	slog.Warn("Server exited")
}
