# dauction

[![GoDoc](https://godoc.org/github.com/kchristidis/dauction?status.svg)](https://godoc.org/github.com/kchristidis/dauction)
[![Build Status](https://travis-ci.org/kchristidis/dauction.svg?branch=master)](https://travis-ci.org/kchristidis/dauction)

dauction faciliates the calculation of double auction clearing prices.

## Installation

```bash
$ go get github.com/kchristidis/dauction
```

## Usage

```go
// group buyer bids into a bid collection object
bb1 := dauction.Bid{PricePerUnit: 6.5, Units: 2}
bb2 := dauction.Bid{PricePerUnit: 10, Units: 2}
buyerBids := dauction.BidCollection{bb1, bb2}

// same for seller bids
sb1 := dauction.Bid{PricePerUnit: 6.5, Units: 2}
sb2 := dauction.Bid{PricePerUnit: 11, Units: 2}
sellerBids := dauction.BidCollection{sb1, sb2}

// settle the market
res, err := dauction.Settle(buyerBids, sellerBids)
if err != nil { // When no clearing price can be found
    fmt.Println(err)
}
// - clearing price: res.PricePerUnit
// - number of units that can be cleared: res.Units
fmt.Println(res)
```

You may also wish to consult the package documentation in [GoDoc](http://godoc.org/github.com/kchristidis/overlap).

## Contributing

Contributions are welcome. Fork this library and submit a pull request.