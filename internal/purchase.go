package internal

import (
	"errors"
	"time"
)

type Purchase struct {
	Id     int       `json:"id"`
	Vendor string    `json:"vendor"`
	Name   string    `json:"name"`
	Shop   string    `json:"shop"`
	Price  float64   `json:"price"`
	Qty    int       `json:"qty"`
	Time   time.Time `json:"time"`
}

func NewPurchase(id int, vendor string, name string, shop string, price float64, qty int, time time.Time) (Purchase, error) {
	if price < 0.0 {
		return Purchase{}, errors.New("price can't be negative")
	}
	if qty <= 0 {
		return Purchase{}, errors.New("qty can't be 0 or negative")
	}

	purchase := Purchase{
		Id:     id,
		Vendor: vendor,
		Name:   name,
		Shop:   shop,
		Price:  price,
		Qty:    qty,
		Time:   time,
	}

	return purchase, nil
}
