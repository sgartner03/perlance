package internal

import (
	"errors"
	"time"
)

type Purchase struct {
	id     int       `json:"id"`
	vendor string    `json:"vendor"`
	name   string    `json:"title"`
	shop   string    `json:"title"`
	price  float64   `json:"title"`
	qty    int       `json:"title"`
	time   time.Time `json:"title"`
}

func NewPurchase(id int, vendor string, name string, shop string, price float64, qty int, time time.Time) (Purchase, error) {
	if price < 0.0 {
		return Purchase{}, errors.New("price can't be negative")
	}
	if qty <= 0 {
		return Purchase{}, errors.New("qty can't be 0 or negative")
	}

	purchase := Purchase{
		id:     id,
		vendor: vendor,
		name:   name,
		shop:   shop,
		price:  price,
		qty:    qty,
		time:   time,
	}

	return purchase, nil
}

func (p *Purchase) GetId() int {
	return p.id
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
func (p *Purchase) GetPrice() float64 {
	return p.price
}
func (p *Purchase) GetQty() int {
	return p.qty
}
func (p *Purchase) GetTime() time.Time {
	return p.time
}
