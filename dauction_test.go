package dauction_test

import (
	"testing"

	"github.com/kchristidis/dauction"
	"github.com/stretchr/testify/assert"
)

const epsilon = 1E-6

func TestBilateralTrade(t *testing.T) {
	buyers := []dauction.Bid{
		dauction.Bid{Units: 1, PricePerUnit: 10},
	}
	sellers := []dauction.Bid{
		dauction.Bid{Units: 1, PricePerUnit: 8},
	}
	expected := float32(9)
	actual, err := dauction.Price(buyers, sellers)
	assert.NoError(t, err, "Expected a clearing price to be found")
	assert.InEpsilon(t, expected, actual, epsilon,
		"Expected a trading price of %.2f, got %.2f instead", expected, actual)
}
