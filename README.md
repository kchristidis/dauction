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
var buyerBids, sellerBids []dauction.Bid
buyerBids = append(buyerBids, dauction.Bid{1,10})
sellerBids = append(sellerBids, dauction.Bid{1, 8})
price, err := dauction.Price(buyerBids, sellerBids)
if err != nil { // When no clearing price can be found
    println(err)
}
println("The clearing price is:", price)
```

You may also wish to consult the package documentation in [GoDoc](http://godoc.org/github.com/kchristidis/overlap).

## Contributing

Contributions are welcome. Fork this library and submit a pull request.