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

func NewPurchase(vendor string, name string, shop string, price float32, qty int, time time.Time) (*Purchase, error) {
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

func (p *Purchase) GetVendor() string {
	return p.vendor
}

func (p *Purchase) GetName() string {
	return p.name
}
func (p *Purchase) GetShop() string {
	return p.shop
}
func (p *Purchase) GetPrice() float32 {
	return p.price
}
func (p *Purchase) GetQty() int {
	return p.qty
}
func (p *Purchase) GetTime() time.Time {
	return p.time
}
