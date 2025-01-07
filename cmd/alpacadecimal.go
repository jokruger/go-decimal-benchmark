package main

import (
	decimal "github.com/alpacahq/alpacadecimal"
)

type alpacadecimalPair struct {
	a decimal.Decimal
	b decimal.Decimal
}

type AlpacadecimalTester struct {
	parseCases []string
	addCases   []alpacadecimalPair
	mulCases   []alpacadecimalPair
	divCases   []alpacadecimalPair
}

func NewAlpacadecimalTester() Tester {
	return &AlpacadecimalTester{}
}

func (self *AlpacadecimalTester) Name() string {
	return "alpacadecimal.Decimal"
}

func (self *AlpacadecimalTester) Init() {
	self.parseCases = ParseCases

	parsePairs := func(pairs []Pair) []alpacadecimalPair {
		result := make([]alpacadecimalPair, len(pairs))
		for i, p := range pairs {
			a, err := decimal.NewFromString(p.A)
			if err != nil {
				panic(err)
			}
			b, err := decimal.NewFromString(p.B)
			if err != nil {
				panic(err)
			}
			result[i] = alpacadecimalPair{a, b}
		}
		return result
	}

	self.addCases = parsePairs(AddCases)
	self.mulCases = parsePairs(MulCases)
	self.divCases = parsePairs(DivCases)
}

func (self *AlpacadecimalTester) RunParse() int64 {
	var n int64
	for range 100000 {
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

func (self *AlpacadecimalTester) RunString() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.addCases {
			_ = c.a.String()
			_ = c.b.String()
			n += 2
		}
	}
	return n
}

func (self *AlpacadecimalTester) RunAdd() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.addCases {
			_ = c.a.Add(c.b)
			n++
		}
	}
	return n
}

func (self *AlpacadecimalTester) RunMul() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.mulCases {
			_ = c.a.Mul(c.b)
			n++
		}
	}
	return n
}

func (self *AlpacadecimalTester) RunDiv() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.divCases {
			_ = c.a.Div(c.b)
			n++
		}
	}
	return n
}
