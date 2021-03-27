package pricing

import "fmt"

var currentInventory = map[string]float64{
	"A": 50, "B": 30, "C": 20, "D": 15,
}

type offer struct {
	offerID   int
	sku       string
	quantity  int
	discount  float64
	offerDesc string
}

// Note:
// 1. currentOffers are initialized with a lower case(which means it is abstracted / not exported)
// so the checkout service won't be aware of pricing strategies (as mentioned in the problem statement)
// 2. A separate function can be written to add and delete offers in this list
// so that we can add new pricing rules in the future without changing the code (as mentioned in the problem statement)
// current offers - {id, "sku", quantity, discount, "offer description"}
var currentOffers = []offer{
	{1, "A", 3, 20, "3A for 130"},
	{2, "B", 2, 15, "2B for 45"},
	{3, "C", 5, 20, "5C for 80"},
}

// ValidateSKU function to validate if input SKU is valid or not
func ValidateSKU(sku string) bool {
	var res bool

	_, res = currentInventory[sku]

	// return
	return res
}

// GetCartPrice to get final cart price
func GetCartPrice(cart map[string]int) (totalCartValue, discount, finalCartValue float64, offersApplied []string) {

	// calculate discount
	for item, quantity := range cart {
		for _, v := range currentOffers {
			if v.sku == item && quantity >= v.quantity {
				discount += float64(quantity/v.quantity) * v.discount
				offersApplied = append(offersApplied, fmt.Sprintf("%d", quantity/v.quantity)+"* "+v.offerDesc)
			}
		}
	}

	// calculate total cart value
	for item, quantity := range cart {
		totalCartValue += currentInventory[item] * float64(quantity)
	}

	// Final Cart Value
	finalCartValue = totalCartValue - discount

	// return
	return totalCartValue, discount, finalCartValue, offersApplied
}
