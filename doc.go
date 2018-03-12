/*
Package dauction facilitates the calculation of double auction clearing prices.

Define a double auction

First group the buyer-submitted bids into a slice of type Bid. Each Bid carries
the number of units a buyer wishes to buy, as well as the maximum price they
are willing to pay. This price is per unit and applies to all units in their
Bid.

Then proceed similary for the seller-submitted offers, grouping them into a
slice of type Bid. Each Bid carries the number of units a seller wishes to
sell, as well as the minimum price they demand per unit.

Determine the clearing price

Pass the two slices to the Price function. The price is calculated using the
average mechanism.

	price, err := Price(buyerBids, sellerBids)

An error will be returned if no clearing (equilibrium) price can be found for
that market.
*/
package dauction
