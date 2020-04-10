/*
Package dauction facilitates the calculation of double auction clearing prices.

Define a double auction

First group the buyer-submitted bids into a BidCollection object. Each Bid
carries the maximum number of units a buyer wishes to buy, as well as the
maximum price they are willing to pay. This price is per unit and applies to
all units in their Bid.

Then proceed similary for the seller-submitted offers, grouping them into a
BidCollection object. Each Bid carries the maximum number of units a seller
wishes to sell, as well as the minimum price they demand per unit.

Determine the clearing price

Pass the two bid collections to the Settle function.

	res, err := dauction.Settle(buyerBids, sellerBids)

The result carries the market-clearing price (res.PricePerUnit) and the number
of units that can be traded (res.Units) given that price.

The price is calculated using the average mechanism, with a bias towards
economic efficiency. Simply put, the calculated market-clearing price allows
for the largest possible number of units to be traded. When the number of
tradeable units is maximized for more than one price point, the average value
across all candidate price points is chosen as the market clear price.

Settle will return an error if no clearing (equilibrium) price can be found for
that market.
*/
package dauction
