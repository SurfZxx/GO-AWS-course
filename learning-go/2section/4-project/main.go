package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    5.00,
	"BOOK":   15.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf(" - Item %s (Sale! Original: $%.2f, Sale Price: $%.2f)\n",
					originalItemCode,
					basePrice,
					salePrice,
				)
				return salePrice, true
			}
		}
		fmt.Println("Product not found in product prices")
		return 0.0, false
	}
	return basePrice, true
}

func main() {
	fmt.Println("---- Simple Sales Order Process")

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	var subTotal float64
	fmt.Println("--- Processing Order Items:")
	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subTotal += price
		}
	}
	fmt.Println("Total Price:", subTotal)
}
