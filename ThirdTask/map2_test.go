package ThirdTask

import (
	"errors"
	"github.com/segmentio/fasthash/fnv1a"
	"math"
	"reflect"
	"testing"
)

func hashOrder(order []string) uint64 {
	m := make(map[string]float32)
	for i := range order {
		m[order[i]] = 0
	}
	order = sortMapByName(&m)

	var hash string
	for _, item := range order {
		hash += item
	}

	return fnv1a.HashString64(hash)
}

func TestSortByName(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			"basic case. nil slice",
			nil,
			nil,
		},
		{
			"basic case. empty slice",
			[]string{},
			[]string{},
		},
		{
			"basic case. single element slice",
			[]string{"a"},
			[]string{"a"},
		},
		{
			"single duplicate element. same price",
			[]string{"a", "a"},
			[]string{"a", "a"},
		},
		{
			"single duplicate element with one in the middle. different price",
			[]string{"a", "b", "a"},
			[]string{"a", "a", "b"},
		},
		{
			"five elements with duplicates",
			[]string{"a", "b", "c", "a", "c"},
			[]string{"a", "a", "b", "c", "c"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := make(map[string]float32)
			for i := range tc.input {
				m[tc.input[i]] = 0
			}
			res := sortMapByName(&m)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}

func TestHashOrder(t *testing.T) {
	//no func
}

type Item struct {
	name  string
	price uint64
}

type shopError string

func (err shopError) Error() string {
	return string(err)
}

const (
	errItemNoFound          shopError = "the item not found"
	errEmptyItem            shopError = "an empty item"
	errItemAlreadyExists    shopError = "the item already exists"
	errEmptyAccount         shopError = "an empty account"
	errAccountNoFound       shopError = "the account not found"
	errAccountAlreadyExists shopError = "the account already exists"
	errAccountUnknownSort   shopError = "an unknown sort"
)

func TestCalculateOrder(t *testing.T) {
	testCases := []struct {
		name               string
		shop               map[string]Item
		order              []string
		expectedTotalPrice uint64
		expectedError      error
	}{
		{
			"basic case. nil slice",
			map[string]Item{},
			nil,
			0,
			nil,
		},
		{
			"basic case. empty slice",
			map[string]Item{},
			[]string{},
			0,
			nil,
		},
		{
			"basic case. single element slice",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"a"},
			1,
			nil,
		},
		{
			"basic case. two element slice",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"b", "a"},
			11,
			nil,
		},
		{
			"basic case. single unknown item",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"xxx"},
			0,
			errItemNoFound,
		},
		{
			"basic case. single unknown item inbetween ",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"a", "xxx", "b"},
			0,
			errItemNoFound,
		},
		{
			"partial match",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"aa"},
			0,
			errItemNoFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			myShop := make(map[string]float32)

			for i := range tc.shop {
				myShop[i] = float32(tc.shop[i].price)
			}

			res := uint64(CalculateSumOrder(&myShop, tc.order))
			var err error

			if !errors.Is(err, tc.expectedError) {
				t.Fatalf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(res, tc.expectedTotalPrice) {
				t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expectedTotalPrice)
			}
		})
	}
}

func TestCalculateOrderWithCache(t *testing.T) {
	testCases := []struct {
		name               string
		shop               map[string]Item
		order              []string
		expectedTotalPrice uint64
		expectedError      error
	}{
		{
			"basic case. nil slice",
			map[string]Item{},
			nil,
			0,
			nil,
		},
		{
			"basic case. empty slice",
			map[string]Item{},
			[]string{},
			0,
			nil,
		},
		{
			"basic case. single element slice",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"a"},
			1,
			nil,
		},
		{
			"basic case. two element slice",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"b", "a"},
			11,
			nil,
		},
		{
			"basic case. single unknown item",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"xxx"},
			0,
			errItemNoFound,
		},
		{
			"basic case. single unknown item inbetween ",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"a", "xxx", "b"},
			0,
			errItemNoFound,
		},
		{
			"partial match",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			[]string{"aa"},
			0,
			errItemNoFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cache := make(map[uint64]uint64)
			for i := 0; i < 5; i++ {
				myShop := make(map[string]float32)

				for i := range tc.shop {
					myShop[i] = float32(tc.shop[i].price)
				}
				fun := CalculateSumOrderWithMemory()

				res := uint64(fun(&myShop, tc.order))

				var err error

				if !errors.Is(err, tc.expectedError) {
					t.Fatalf("got\t\t%v\nwant\t%v", err, tc.expectedError)
				}
				if tc.expectedError == nil && i > 0 {
					totalFromCache, ok := cache[hashOrder(tc.order)]
					if !ok {
						t.Fatalf("cache was not used: %v", tc.order)
					}
					if totalFromCache != tc.expectedTotalPrice {
						t.Fatalf("got in cache\t\t%v\nwant\t%v", totalFromCache, tc.expectedTotalPrice)
					}
				}
				if !reflect.DeepEqual(res, tc.expectedTotalPrice) {
					t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expectedTotalPrice)
				}
			}
		})
	}
}

func TestAddItem(t *testing.T) {
	testCases := []struct {
		name          string
		shop          map[string]Item
		item          Item
		expectedShop  map[string]Item
		expectedError error
	}{
		{
			"basic case. empty Item",
			map[string]Item{},
			Item{},
			map[string]Item{},
			errEmptyItem,
		},
		{
			"basic case. empty item name",
			map[string]Item{},
			Item{"", 10},
			map[string]Item{},
			errEmptyItem,
		},
		{
			"basic case. already exists",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Item{"a", 10},
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			errItemAlreadyExists,
		},
		{
			"basic case. correct item",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Item{"xxx", 100},
			map[string]Item{
				"a":   {"a", 1},
				"b":   {"b", 10},
				"xxx": {"xxx", 100},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			myShop := make(map[string]float32)

			for i := range tc.shop {
				myShop[i] = float32(tc.shop[i].price)
			}

			AddProduct(&myShop, tc.item.name, float32(tc.item.price))

			var err error

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(tc.shop, tc.expectedShop) {
				t.Fatalf("got\t\t%v\nwant\t%v", tc.shop, tc.expectedShop)
			}
		})
	}
}

func TestChangePrice(t *testing.T) {
	testCases := []struct {
		name          string
		shop          map[string]Item
		item          Item
		expectedShop  map[string]Item
		expectedError error
	}{
		{
			"basic case. empty Item",
			map[string]Item{},
			Item{},
			map[string]Item{},
			errEmptyItem,
		},
		{
			"basic case. empty item name",
			map[string]Item{},
			Item{"", 10},
			map[string]Item{},
			errEmptyItem,
		},
		{
			"basic case. correct price change",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Item{"a", 10},
			map[string]Item{
				"a": {"a", 10},
				"b": {"b", 10},
			},
			nil,
		},
		{
			"basic case. correct price change",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Item{"xxx", 10},
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			errItemNoFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			myShop := make(map[string]float32)

			for i := range tc.shop {
				myShop[i] = float32(tc.shop[i].price)
			}
			ChangePriceProduct(&myShop, tc.item.name, float32(tc.item.price))

			var err error

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(tc.shop, tc.expectedShop) {
				t.Fatalf("got\t\t%v\nwant\t%v", tc.shop, tc.expectedShop)
			}
		})
	}
}

func TestChangeName(t *testing.T) {
	testCases := []struct {
		name          string
		shop          map[string]Item
		itemName      string
		newItemName   string
		expectedShop  map[string]Item
		expectedError error
	}{
		{
			"basic case. empty Item",
			map[string]Item{},
			"a",
			"",
			map[string]Item{},
			errEmptyItem,
		},
		{
			"basic case. empty item name",
			map[string]Item{
				"a": {"a", 10},
			},
			"a",
			"",
			map[string]Item{
				"a": {"a", 10},
			},
			errEmptyItem,
		},
		{
			"basic case. already exists",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			"a",
			"aa",
			map[string]Item{
				"aa": {"aa", 1},
				"b":  {"b", 10},
			},
			nil,
		},
		{
			"basic case. already exists",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			"xxx",
			"aa",
			map[string]Item{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			errItemNoFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			myShop := make(map[string]float32)

			for i := range tc.shop {
				myShop[i] = float32(tc.shop[i].price)
			}
			ChangeNameProduct(&myShop, tc.itemName, tc.newItemName)

			var err error

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(tc.shop, tc.expectedShop) {
				t.Fatalf("got\t\t%v\nwant\t%v", tc.shop, tc.expectedShop)
			}
		})
	}
}

func TestAddAccount(t *testing.T) {
	testCases := []struct {
		name             string
		accounts         map[string]Account
		account          Account
		expectedAccounts map[string]Account
		expectedError    error
	}{
		{
			"basic case. empty Account",
			map[string]Account{},
			Account{},
			map[string]Account{},
			errEmptyAccount,
		},
		{
			"basic case. empty Account name",
			map[string]Account{},
			Account{"", 10},
			map[string]Account{},
			errEmptyAccount,
		},
		{
			"basic case. already exists",
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Account{"a", 10},
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			errAccountAlreadyExists,
		},
		{
			"basic case. correct item",
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Account{"xxx", 100},
			map[string]Account{
				"a":   {"a", 1},
				"b":   {"b", 10},
				"xxx": {"xxx", 100},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			myAccounts := make(map[string]float32)
			for i := range myAccounts {
				myAccounts[i] = float32(tc.accounts[i].balance)
			}

			AddAccount(&myAccounts, tc.account.name, float32(tc.account.balance))

			var err error

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(tc.accounts, tc.expectedAccounts) {
				t.Fatalf("got\t\t%v\nwant\t%v", tc.account, tc.expectedAccounts)
			}
		})
	}
}

func TestChangeBalance(t *testing.T) {
	// no func
}

type SortBy uint8

const (
	SortByNameAsc     SortBy = iota
	SortByNameDesc    SortBy = iota
	SortByBalanceAsc  SortBy = iota
	SortByBalanceDesc SortBy = iota
)

func TestSortAccounts(t *testing.T) {
	testCases := []struct {
		name             string
		accounts         map[string]Account
		sortBy           SortBy
		expectedAccounts []Account
		expectedError    error
	}{
		// name
		{
			"SortByNameAsc. empty Accounts",
			map[string]Account{},
			SortByNameAsc,
			[]Account{},
			nil,
		},
		{
			"SortByNameAsc. single Account ",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByNameAsc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByNameAsc. already sorted",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"xxx_1": {"xxx_1", 22},
			},
			SortByNameAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d", 12},
				{"d10", 11},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByNameAsc. already sorted in reverse order",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByNameAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d", 12},
				{"d10", 11},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByNameAsc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByNameAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d", 12},
				{"d10", 11},
				{"xxx_1", 22},
			},
			nil,
		},

		{
			"SortByNameDesc. empty Accounts",
			map[string]Account{},
			SortByNameDesc,
			[]Account{},
			nil,
		},
		{
			"SortByNameDesc. single Account ",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByNameDesc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByNameDesc. already sorted",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByNameDesc,
			[]Account{
				{"xxx_1", 22},
				{"d10", 11},
				{"d", 12},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByNameDesc. already sorted in reverse order",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"xxx_1": {"xxx_1", 22},
			},
			SortByNameDesc,
			[]Account{
				{"xxx_1", 22},
				{"d10", 11},
				{"d", 12},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByNameDesc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByNameDesc,
			[]Account{
				{"xxx_1", 22},
				{"d10", 11},
				{"d", 12},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},

		// balance
		{
			"SortByBalanceAsc. empty Accounts",
			map[string]Account{},
			SortByBalanceAsc,
			[]Account{},
			nil,
		},
		{
			"SortByBalanceAsc. single Account",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByBalanceAsc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted with duplicates",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d11":   {"d11", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d11", 11},
				{"d", 12},
				{"xxx_1", 22},
				{"xxx_2", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted in reverse order",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted in reverse order with duplicated",
			map[string]Account{
				"xxx_2": {"xxx_2", 22},
				"xxx_1": {"xxx_1", 22},
				"d":     {"d", 12},
				"d11":   {"d11", 11},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d11", 11},
				{"d", 12},
				{"xxx_1", 22},
				{"xxx_2", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. random order with duplicates",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"a1":    {"a1", 1},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"a1", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
				{"xxx_2", 22},
			},
			nil,
		},

		{
			"SortByBalanceDesc. empty Accounts",
			map[string]Account{},
			SortByBalanceDesc,
			[]Account{},
			nil,
		},
		{
			"SortByBalanceDesc. single Account",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByBalanceDesc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted with duplicates",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
				"d":     {"d", 12},
				"d11":   {"d11", 11},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"xxx_2", 22},
				{"d", 12},
				{"d10", 11},
				{"d11", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted in reverse order",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted in reverse order with duplicated",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d11":   {"d11", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"xxx_2", 22},
				{"d", 12},
				{"d10", 11},
				{"d11", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. random order with duplicates",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"a1":    {"a1", 1},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"xxx_2", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
				{"a1", 1},
			},
			nil,
		},

		// unknown
		{
			"Unknown order on empty accounts",
			map[string]Account{},
			SortBy(math.MaxUint8),
			[]Account{},
			nil,
		},
		{
			"Unknown order error.",
			map[string]Account{
				"a": {"a", 1},
			},
			SortBy(math.MaxUint8),
			nil,
			errAccountUnknownSort,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			myAccounts := make(map[string]float32)
			for i := range tc.accounts {
				myAccounts[i] = float32(tc.accounts[i].balance)
			}

			var res []Account

			switch tc.sortBy {
			case SortByNameAsc:
				res = sortMapByNameA(&myAccounts)
			case SortByNameDesc:
				res = sortMapByNameReverseA(&myAccounts)
			case SortByBalanceDesc:
				res = sortMapByPriceA(&myAccounts)
			default:
				//no func
			}

			var err error

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(res, tc.expectedAccounts) {
				t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expectedAccounts)
			}
		})
	}
}
