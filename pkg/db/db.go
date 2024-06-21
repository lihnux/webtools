package db

import (
	"log/slog"
	"sync"

	"github.com/tidwall/buntdb"
)

var (
	buntdbOnce sync.Once
	buntdbConn *buntdb.DB
)

func BuntDB() *buntdb.DB {
	buntdbOnce.Do(func() {
		var err error
		buntdbConn, err = buntdb.Open("runtime.db")
		if err != nil {
			slog.Error("DB> Cannot open buntdb file 'runtime.db'", "reason", err.Error())
			return
		}

		var config buntdb.Config
		if err := buntdbConn.ReadConfig(&config); err != nil {
			slog.Error("DB> Read buntdb config failed", "reason", err.Error())
			return
		}

		slog.Info("DB> buntdb config", "config", config)

		err = buntdbConn.CreateIndex("stocks", "stock:*", buntdb.IndexString)
		if err != nil {
			slog.Info("DB> create 'stocks' index failed", "reason", err.Error())
		}

		buntdbConn.Shrink()
	})

	return buntdbConn
}

func Disconnect() {
	if buntdbConn != nil {
		buntdbConn.Close()
		buntdbConn = nil
	}
}
