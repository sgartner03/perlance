package internal

import (
	"errors"
	"time"
)

type Purchase struct {
	vendor string
	name   string
	shop   string
	price  float32
	qty    int
	time   time.Time
}

func (*Purchase) Purchase(vendor string, name string, shop string, price float32, qty int, time time.Time) (*Purchase, error) {
	if price < 0.0 {
		return nil, errors.New("price can't be negative")
	}
	if qty <= 0 {
		return nil, errors.New("qty can't be 0 or negative")
	}

	purchase := &Purchase{
		vendor: vendor,
		name:   name,
		shop:   shop,
		price:  price,
		qty:    qty,
		time:   time,
	}

	return purchase, nil
}
