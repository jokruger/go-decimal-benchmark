package main

import (
	"github.com/shopspring/decimal"
)

type shopspringPair struct {
	a decimal.Decimal
	b decimal.Decimal
}

type ShopspringTester struct {
	parseCases []string
	addCases   []shopspringPair
	mulCases   []shopspringPair
	divCases   []shopspringPair
}

func NewShopspringTester() Tester {
	return &ShopspringTester{}
}

func (self *ShopspringTester) Name() string {
	return "shopspring.Decimal"
}

func (self *ShopspringTester) Init() {
	self.parseCases = ParseCases

	parsePairs := func(pairs []Pair) []shopspringPair {
		result := make([]shopspringPair, len(pairs))
		for i, p := range pairs {
			a, err := decimal.NewFromString(p.A)
			if err != nil {
				panic(err)
			}
			b, err := decimal.NewFromString(p.B)
			if err != nil {
				panic(err)
			}
			result[i] = shopspringPair{a, b}
		}
		return result
	}

	self.addCases = parsePairs(AddCases)
	self.mulCases = parsePairs(MulCases)
	self.divCases = parsePairs(DivCases)
}

func (self *ShopspringTester) RunParse() int64 {
	var n int64
	for range 10000 {
		for _, c := range self.parseCases {
			_, err := decimal.NewFromString(c)
			if err != nil {
				panic(err)
			}
			n++
		}
	}
	return n
}

func (self *ShopspringTester) RunAdd() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.addCases {
			_ = c.a.Add(c.b)
			n++
		}
	}
	return n
}

func (self *ShopspringTester) RunMul() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.mulCases {
			_ = c.a.Mul(c.b)
			n++
		}
	}
	return n
}

func (self *ShopspringTester) RunDiv() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.divCases {
			_ = c.a.Div(c.b)
			n++
		}
	}
	return n
}
