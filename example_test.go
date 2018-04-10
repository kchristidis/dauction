package dauction_test

import (
	"fmt"

	"github.com/kchristidis/dauction"
)

func ExampleSettle() {
	// group buyer bids into a bid collection object
	bb1 := dauction.Bid{PricePerUnit: 8, Units: 2}
	bb2 := dauction.Bid{PricePerUnit: 10, Units: 2}
	buyerBids := dauction.BidCollection{bb1, bb2}

	// same for seller bids
	sb1 := dauction.Bid{PricePerUnit: 6, Units: 1}
	sb2 := dauction.Bid{PricePerUnit: 9, Units: 1}
	sellerBids := dauction.BidCollection{sb1, sb2}

	// have the auctioneer clear the market
	res, err := dauction.Settle(buyerBids, sellerBids)
	if err != nil { // when no clearing price can be found
		fmt.Println(err)
	}
	// - res.PricePerUnit = 9.5 (clearing price)
	// - res.Units = 2 (number of units that can be cleared)
	fmt.Println(res)
}
