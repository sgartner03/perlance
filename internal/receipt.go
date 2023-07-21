package internal

import (
	"time"
)

type Receipt struct {
	time      time.Time
	shop      string
	purchases []Purchase
}

func NewReceipt(time time.Time, shop string) (Receipt, error) {
	receipt := Receipt{
		time: time,
		shop: shop,
	}

	return receipt, nil
}

func (receipt *Receipt) Add(id int, vendor string, name string, price float64, qty int) error {
	purchase, err := NewPurchase(id, vendor, name, receipt.shop, price, qty, receipt.time)
	if err != nil {
		return err
	} else {
		receipt.purchases = append(receipt.purchases, purchase)
		return nil
	}
}

func (receipt *Receipt) Get(index int) Purchase {
	return receipt.purchases[index]
}
func (receipt *Receipt) Size() int {
	return len(receipt.purchases)
}

func (receipt *Receipt) GetTime() time.Time {
	return receipt.time
}
func (receipt *Receipt) GetShop() string {
	return receipt.shop
}
func (receipt *Receipt) GetPurchases() []Purchase {
	return receipt.purchases
}
