package model

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"webtools/pkg/db"
	"webtools/pkg/utils"

	"github.com/tidwall/buntdb"
)

const (
	BuyStock = iota
	SellStock
)

type Stock struct {
	Name        string
	Symbol      string
	Description string
}

type Trade struct {
	ID        utils.UniqueID
	Symbol    string
	Operation int32
	Price     float64
	Shares    int32
}

func NewStock(name, symbol, description string) *Stock {
	return &Stock{
		Name:        name,
		Symbol:      symbol,
		Description: description,
	}
}

func ListStocks() []*Stock {
	stocks := []*Stock{}
	db.BuntDB().View(func(tx *buntdb.Tx) error {
		tx.Ascend("stocks", func(key, val string) bool {
			slog.Info("MODEL-STOCK> Found stock", "key", key)
			var stock Stock
			err := json.Unmarshal([]byte(val), &stock)
			if err == nil {
				stocks = append(stocks, &stock)
			}
			return true
		})
		return nil
	})

	return stocks
}

func (s *Stock) String() string {
	data, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(data)
}

func (s *Stock) Save() error {
	return db.BuntDB().Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(fmt.Sprintf("stock:%s", s.Name), s.String(), nil)
		return err
	})
}
