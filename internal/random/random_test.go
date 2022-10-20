package random

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestRandomInt(t *testing.T) {
	tests := []struct {
		name string
		min  int64
		max  int64
	}{
		{
			min:  0,
			max:  100,
			name: "Random number is between 0 and 100",
		},
		{
			min:  1,
			max:  2,
			name: "Random number is between 1 and 2",
		},
		{
			min:  0,
			max:  0,
			name: "Random number equals 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			randInt := RandomInt(test.min, test.max)
			require.GreaterOrEqual(t, randInt, test.min)
			require.LessOrEqual(t, randInt, test.max)
		})
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name      string
		randomStr string
	}{
		{
			name:      "Random string consists alphabetic characters",
			randomStr: RandomString(5),
		},
	}

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.True(t, isAlpha(test.randomStr))
			require.Equal(t, len(test.randomStr), 5)
		})
	}
}

func TestRandomOwner(t *testing.T) {
	tests := []struct {
		name  string
		owner string
	}{
		{
			name:  "Random owner is an alphabetic string, six characters in length",
			owner: RandomOwner(),
		},
	}

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.True(t, isAlpha(test.owner))
			require.Equal(t, len(test.owner), 6)
		})
	}
}

func TestRandomMoney(t *testing.T) {
	tests := []struct {
		name   string
		amount int64
	}{
		{
			name:   "Amount is between 0 and 1000",
			amount: RandomMoney(),
		},
	}

	var min int64 = 0
	var max int64 = 1000

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.GreaterOrEqual(t, test.amount, min)
			require.LessOrEqual(t, test.amount, max)
		})
	}
}

func TestRandomCurrency(t *testing.T) {
	tests := []struct {
		name     string
		currency string
	}{
		{
			name:     "Currency is USD, EUR or CAD",
			currency: RandomCurrency(),
		},
	}

	currencies := []string{"USD", "EUR", "CAD"}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.True(t, contains(currencies, test.currency))
		})
	}
}
