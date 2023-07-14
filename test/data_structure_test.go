package test

import "testing"

func TestCreateValidPurchase(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := 1.59
	qty := 2
	purchase, err := internal.Purchase(vendor, name, shop, price, qty)
	if err != nil {
		t.Errorf("Purchase creation failed, except it should: %s", err)
	}

	if purchase.vendor != vendor {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", vendor, purchase.vendor)
	}

	if purchase.name != name {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", name, purchase.name)
	}
	if purchase.shop != shop {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", vendor, purchase.shop)
	}
	if purchase.price != price {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", vendor, purchase.price)
	}
	if purchase.qty != qty {
		t.Errorf("Inconsistent vendor; Input: %s, Output: %s", vendor, purchase.qty)
	}
}

func TestCreatePurchaseNegativePrice(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := -1.59
	qty := 2
	_, err := internal.Purchase(vendor, name, shop, price, qty)
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreatePurchaseNegativeQty(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := -1.59
	qty := -1
	_, err := internal.Purchase(vendor, name, shop, price, qty)
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}

func TestCreatePurchaseZeroQty(t *testing.T) {
	vendor := "Red Bull"
	name := "Organics Cola"
	shop := "Interspar"
	price := -1.59
	qty := 0
	_, err := internal.Purchase(vendor, name, shop, price, qty)
	if err == nil {
		t.Errorf("Purchase creation sucessful, except it shouldn't")
	}
}
