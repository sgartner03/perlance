package test

import (
	"testing"
	"time"
)

func TestCreateValidPurchase(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := 1.59
	qty := 2
	time := time.Now()
	purchase, err := internal.Purchase(vendor, name, shop, price, qty, time)
	if err != nil {
		t.Errorf("Purchase creation failed, except it should: %s", err)
	}

	if purchase.GetVendor() != vendor {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", vendor, purchase.vendor)
	}

	if purchase.GetName() != name {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", name, purchase.name)
	}
	if purchase.GetShop() != shop {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", vendor, purchase.shop)
	}
	if purchase.GetPrice() != price {
		t.Errorf("Inconsistent price; Input: %s, Output: %s", price, purchase.price)
	}
	if purchase.GetQty() != qty {
		t.Errorf("Inconsistent qty; Input: %s, Output: %s", qty, purchase.qty)
	}
	if purchase.GetTime() != time {
		t.Errorf("Inconsistent time; Input: %s, Output: %s", time, purchase.time)
	}
}

func TestCreatePurchaseNegativePrice(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := -1.59
	qty := 2
	_, err := internal.Purchase(time.Now(), vendor, name, shop, price, qty)
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreatePurchaseNegativeQty(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := 1.59
	qty := -1
	_, err := internal.Purchase(time.Now(), vendor, name, shop, price, qty)
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreatePurchaseZeroQty(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := 1.59
	qty := 0
	_, err := internal.Purchase(time.Now(), vendor, name, shop, price, qty)
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreateValidEmptyReceipt(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, err := internal.Receipt(currentTime, shop)
	if err != nil {
		t.Errorf("Receipt creation failed, except it should: %s", err)
	}
}

func TestFillReceipt(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, _ := internal.Receipt(currentTime, shop)

	receipt.Add("Red Bull", "Organics Cola", 1.59, 3)
	receipt.Add("Innocent", "Orange Smoothie", 3.09, 1)

	if receipt.Size() != 2 {
		t.Errorf("Invalid purchases size")
	}
}

func TestFillReceiptWithInvalidItem(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, _ := internal.Receipt(currentTime, shop)

	receipt.Add("Red Bull", "Organics Cola", 1.59, 3)
	receipt.Add("Innocent", "Orange Smoothie", -3.09, 1)

	if receipt.Size() != 1 {
		t.Errorf("Invalid purchases size")
	}
}

func TestReceiptPurchaseTime(t *testing.T) {
	currentTime := time.Now()
	shop := "Interspar"
	receipt, _ := internal.Receipt(currentTime, shop)

	receipt.Add("Red Bull", "Organics Cola", 1.59, 3)

	purchaseTime := receipt.Get(0).GetTime()
	if purchaseTime != currentTime {
		t.Errorf("Purchase time %s inconsistent to receipt time %s", purchaseTime, currentTime)
	}
}
