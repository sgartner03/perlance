package test

import (
	"github.com/sgartner03/perlance/internal"
	"testing"
	"time"
)

func TestCreateValidPurchase(t *testing.T) {
	id := 0
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := 1.59
	qty := 2
	time := time.Now()
	purchase, err := internal.NewPurchase(id, vendor, name, shop, price, qty, time)
	if err != nil {
		t.Errorf("Purchase creation failed, except it should: %s", err)
	}

	if purchase.Vendor != vendor {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", vendor, purchase.Vendor)
	}

	if purchase.Name != name {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", name, purchase.Name)
	}
	if purchase.Shop != shop {
		t.Errorf("Inconsistent shop; Input: %s, Output: %s", shop, purchase.Shop)
	}
	if purchase.Price != price {
		t.Errorf("Inconsistent price; Input: %f, Output: %f", price, purchase.Price)
	}
	if purchase.Qty != qty {
		t.Errorf("Inconsistent qty; Input: %d, Output: %d", qty, purchase.Qty)
	}
	if purchase.Time != time {
		t.Errorf("Inconsistent time; Input: %s, Output: %s", time, purchase.Time)
	}
}

func TestCreatePurchaseNegativePrice(t *testing.T) {
	id := 0
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := -1.59
	qty := 2
	_, err := internal.NewPurchase(id, vendor, name, shop, price, qty, time.Now())
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreatePurchaseNegativeQty(t *testing.T) {
	id := 0
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := 1.59
	qty := -1
	_, err := internal.NewPurchase(id, vendor, name, shop, price, qty, time.Now())
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreatePurchaseZeroQty(t *testing.T) {
	id := 0
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := 1.59
	qty := 0
	_, err := internal.NewPurchase(id, vendor, name, shop, price, qty, time.Now())
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreateValidEmptyReceipt(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, err := internal.NewReceipt(currentTime, shop)
	if err != nil {
		t.Errorf("Receipt creation failed, except it should not: %s", err)
	}

	if receipt.GetTime() != currentTime {
		t.Errorf("Inconsistent time; Input: %s, Output: %s", currentTime, receipt.GetTime())
	}
	if receipt.GetShop() != shop {
		t.Errorf("Inconsistent shop; Input: %s, Output: %s", shop, receipt.GetShop())
	}
	if receipt.GetPurchases() != nil {
		t.Errorf("Purchases is initialized")
	}

}

func TestFillReceipt(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, _ := internal.NewReceipt(currentTime, shop)

	receipt.Add(0, "Red Bull", "Organics Cola", 1.59, 3)
	receipt.Add(1, "Innocent", "Orange Smoothie", 3.09, 1)

	if receipt.Size() != 2 {
		t.Errorf("Invalid purchases size: %d", receipt.Size())
	}
}

func TestFillReceiptWithInvalidItem(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, _ := internal.NewReceipt(currentTime, shop)

	receipt.Add(0, "Red Bull", "Organics Cola", 1.59, 3)
	receipt.Add(1, "Innocent", "Orange Smoothie", -3.09, 1)

	if receipt.Size() != 1 {
		t.Errorf("Invalid purchases size: %d", receipt.Size())
	}
}

func TestReceiptPurchaseTime(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, _ := internal.NewReceipt(currentTime, shop)

	receipt.Add(0, "Red Bull", "Organics Cola", 1.59, 3)

	purchase0 := receipt.Get(0)
	purchaseTime := purchase0.Time
	if purchaseTime != currentTime {
		t.Errorf("Purchase time %s inconsistent to receipt time %s", purchaseTime, currentTime)
	}
}

func TestReceiptPurchaseShop(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, _ := internal.NewReceipt(currentTime, shop)

	receipt.Add(0, "Red Bull", "Organics Cola", 1.59, 3)

	receipt0 := receipt.Get(0)
	purchaseShop := receipt0.Shop
	if purchaseShop != shop {
		t.Errorf("Purchase time %s inconsistent to receipt time %s", purchaseShop, shop)
	}
}
