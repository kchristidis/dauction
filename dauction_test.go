package dauction_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/kchristidis/dauction"
	"github.com/stretchr/testify/assert"
)

// The default value for the acceptable relative error
// during clearing price and unit number calculations.
const defaultEpsilon = 1E-5

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestSortBids(t *testing.T) {
	bidPrices := sort.Float64Slice{
		4.20,
		2.05,
		8.14,
	}

	var sellers []dauction.Bid

	for i := range bidPrices {
		sellers = append(sellers, dauction.Bid{PricePerUnit: bidPrices[i]})
	}

	sellersGroup := dauction.BidCollection(sellers)

	// Test in increasing order.
	sort.Sort(bidPrices)
	sort.Sort(sellersGroup)
	for i := range sellers {
		assert.InEpsilon(t, bidPrices[i], sellersGroup[i].PricePerUnit, defaultEpsilon)
	}

	// Test in decreasing order.
	sort.Sort(sort.Reverse(sellersGroup))
	sort.Sort(sort.Reverse(bidPrices))
	for i := range sellers {
		assert.InEpsilon(t, bidPrices[i], sellersGroup[i].PricePerUnit, defaultEpsilon)
	}
}

func TestStackBids(t *testing.T) {
	bids := []dauction.Bid{
		{PricePerUnit: 2.0, Units: 10},
		{PricePerUnit: 1.0, Units: 10},
		{PricePerUnit: 3.0, Units: 10},
	}

	bc := dauction.BidCollection(bids)

	actualResult := dauction.Stack(bc, dauction.Buyers)

	buyersResult := []dauction.Bid{
		{PricePerUnit: 1.0, Units: 30},
		{PricePerUnit: 2.0, Units: 20},
		{PricePerUnit: 3.0, Units: 10},
	}

	for i := range actualResult {
		assert.InDelta(t, buyersResult[i].PricePerUnit, actualResult[i].PricePerUnit, defaultEpsilon,
			fmt.Sprintf("Got %s, but expected %s", actualResult[i], buyersResult[i]))
		assert.InDelta(t, buyersResult[i].Units, actualResult[i].Units, defaultEpsilon,
			fmt.Sprintf("Got %s, but expected %s", actualResult[i], buyersResult[i]))
	}

	sellersResult := []dauction.Bid{
		{PricePerUnit: 3.0, Units: 30},
		{PricePerUnit: 2.0, Units: 20},
		{PricePerUnit: 1.0, Units: 10},
	}

	actualResult = dauction.Stack(bc, dauction.Sellers)

	for i := range actualResult {
		assert.InDelta(t, sellersResult[i].PricePerUnit, actualResult[i].PricePerUnit, defaultEpsilon,
			fmt.Sprintf("Got %s, but expected %s", actualResult[i], sellersResult[i]))
		assert.InDelta(t, sellersResult[i].Units, actualResult[i].Units, defaultEpsilon,
			fmt.Sprintf("Got %s, but expected %s", actualResult[i], sellersResult[i]))
	}
}

func TestTrades(t *testing.T) {
	type testCase struct {
		buyers, sellers dauction.BidCollection
		expected        dauction.Bid
		err             error
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		buyers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 10, Units: 1},
		},
		sellers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 8, Units: 1},
		},
		expected: dauction.Bid{
			PricePerUnit: 9,
			Units:        1,
		},
		err: nil,
	})

	testCases = append(testCases, testCase{
		buyers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 8, Units: 1},
		},
		sellers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 10, Units: 1},
		},
		err: dauction.ErrNoPrice,
	})

	testCases = append(testCases, testCase{
		buyers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 10, Units: 4},
		},
		sellers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 9, Units: 2},
		},
		expected: dauction.Bid{
			PricePerUnit: 9.5,
			Units:        2,
		},
		err: nil,
	})

	testCases = append(testCases, testCase{
		buyers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 6.5, Units: 2},
			dauction.Bid{PricePerUnit: 10, Units: 2},
		},
		sellers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 6, Units: 2},
		},
		expected: dauction.Bid{
			PricePerUnit: 8,
			Units:        2,
		},
		err: nil,
	})

	testCases = append(testCases, testCase{
		buyers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 6.5, Units: 2},
			dauction.Bid{PricePerUnit: 10, Units: 2},
		},
		sellers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 6.5, Units: 2},
			dauction.Bid{PricePerUnit: 11, Units: 2},
		},
		expected: dauction.Bid{
			PricePerUnit: 8.25,
			Units:        2,
		},
		err: nil,
	})

	testCases = append(testCases, testCase{
		buyers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 8, Units: 2},
			dauction.Bid{PricePerUnit: 10, Units: 2},
		},
		sellers: dauction.BidCollection{
			dauction.Bid{PricePerUnit: 6, Units: 1},
			dauction.Bid{PricePerUnit: 9, Units: 1},
		},
		expected: dauction.Bid{
			PricePerUnit: 9.5,
			Units:        2,
		},
		err: nil,
	})

	for _, tc := range testCases {
		actual, err := dauction.Settle(tc.buyers, tc.sellers)
		if tc.err == nil {
			assert.NoError(t, err, "Expected the market to clear")
		} else {
			assert.EqualError(t, tc.err, dauction.ErrNoPrice.Error())
			continue
		}
		assert.InEpsilon(t, tc.expected.PricePerUnit, actual.PricePerUnit, defaultEpsilon,
			"Expected the market to settle at %s, got %s instead", tc.expected, actual)
		assert.InEpsilon(t, tc.expected.Units, actual.Units, defaultEpsilon,
			"Expected the market to settle at %s, got %s instead", tc.expected, actual)
	}
}
