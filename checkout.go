/*
Problem Statement
We have to implement the code for a supermarket checkout that calculates the total price of a number of items.
In a normal supermarket, items are identified using Stock Keeping Units, or SKUs.
To make things simple, we’ll use individual letters of the alphabet (A, B, C, and so on) as SKUs in our store for now.

Our goods are priced individually. In addition, some items are multi priced: Buy n of them, and they’ll cost you y dollars. For example, item ‘A’ might cost 50 dollars(Unit Price) individually, but this week we have a special offer: Buy three ‘A’s and they’ll cost you $130 (Special Price).

Examples of this week’s Prices are:
Item Unit Price Special Price
A      50    3 for 130
B      30    2 for 45
C      20    5 FOR 80
D      15

Our checkout accepts items in any order, so that if we scan a B, an A, and another
B, we’ll recognize the two B’s and price them at 45 (for a total price so far of 95).
An Example interface to the checkout would look something like the following:
var co = new checkOut (pricingService);
co.scan(item);
co.scan(item);
: :
var OrderTotal = co.CalculateMyOrderTotal();

Expectation
We would require you to implement this checkout which calculates the
total price of the checked out items.
We are looking for an object oriented solution for this.
We are looking for automated unit tests to show that we have covered
the scenarios.
Please keep in mind that the description doesn’t mention the format of
the pricing rules. When you design your solution, consider the following
so that you will be able to come up with a solution which is extensible.
1. How can these pricing be specified in a way that the checkout doesn’t
know about particular items and their pricing strategies?
2. How can we make the design flexible enough so that we can add new
pricing rules in the future without changing the code?

*/

package main

import (
	"fmt"
	pricing "supermarket_checkout/lib"
)

type Item struct {
	Items []string `json:"items"`
}

// Global variables
var cart = map[string]int{}
var flagScanCompleted = false

type chkout struct{}

type checkout interface {
	scan(file string)
	// delete() // for future reference: to delete item from cart
}

func (c chkout) scan(file string) {
	/*
		Steps:
			1. Check if items/sku is valid, if not print error
			2. If item/sku is valid,
				2.1 Add item to cart
				2.2 Call "GetCartPrice" from pricing service
	*/
	var s string
	fmt.Scanf("%s", &s)

	if s == "" {
		flagScanCompleted = true
		// Print totalCartValue, discount, finalCartValue, offersApplied
		totalCartValue, discount, finalCartValue, offersApplied := pricing.GetCartPrice(cart)
		fmt.Printf("TotalCartValue:%v,\n Discount:%v,\n FinalCartValue:%v,\n Offers Applied:%+q\n\n", totalCartValue, discount, finalCartValue, offersApplied)

	} else {
		fmt.Println("Item Scanned:", s)
		// validate input sku
		validSKU := pricing.ValidateSKU(s)
		if !validSKU {
			fmt.Println("SKU unidentified. Try again")
			return
		}
		cart[s]++
		totalCartValue, discount, finalCartValue, offersApplied := pricing.GetCartPrice(cart)
		// print totalCartValue, discount, finalCartValue, offersApplied
		fmt.Printf("\nTotalCartValue:%v,\nDiscount:%v,\nFinalCartValue:%v,\nOffers Applied:%+q\n\n", totalCartValue, discount, finalCartValue, offersApplied)
	}
}

func main() {
	co := chkout{}
	fmt.Println("Scan items:")
	for flagScanCompleted == false {
		co.scan("")
	}

}
