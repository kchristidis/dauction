package dauction

import "errors"

// Bid defines the number of units and the maximum/minimum
// per-unit price a buyer/seller is willing to pay/receive.
type Bid struct {
	Units, PricePerUnit float32
}

// Price determines the clearing price given a list of bids
// from buyers and sellers. If no equilibrium price can be
// found, the nil value is returned.
func Price(buyers []Bid, sellers []Bid) (float32, error) {
	return 0, errors.New("cannot find a price that clears the market")
}
