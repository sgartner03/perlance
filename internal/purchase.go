package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"
)

var client influxdb2.Client

func init() {
	token := "_XNTp2Hh7cbuNkjK4fHQBbmpoA2TMcPc6Wxz5FX0oC0Ljzyh897_m2o1t483iGPADgICRzwIjnFteUXMf38vgw==" //os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client = influxdb2.NewClient(url, token)
}

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

func (purchase *Purchase) WritePurchase() {
	org := "perlance-devs"
	bucket := "perlance"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	tags := map[string]string{
		"vendor": purchase.Vendor,
		"name":   purchase.Name,
		"shop":   purchase.Shop,
		"price":  fmt.Sprintf("%f", purchase.Price),
		"qty":    fmt.Sprintf("%d", purchase.Qty),
	}
	fields := map[string]interface{}{
		"amount": float64(purchase.Qty) * purchase.Price,
	}
	point := write.NewPoint(
		"purchase",
		tags, fields, purchase.Time)

	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		log.Fatal(err)
	}
}

type FilterOption func(query *string) error

func WithFilter(key string, value string) FilterOption {
	return func(query *string) error {
		keys := []string{"vendor", "name", "shop", "price", "qty"}

		if !Contains(key, keys) {
			return errors.New("No matching option found")
		}
		*query += fmt.Sprintf("    |> filter(fn: (r) => r.%s == \"%s\"", key, value)
		return nil
	}
}

func Contains(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func (purchase *Purchase) GetPurchases(opts ...FilterOption) {
	query := ""
	for _, option := range opts {
		option(&query)
	}
}
