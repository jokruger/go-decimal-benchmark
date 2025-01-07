package main

import (
	"github.com/quagmt/udecimal"
)

type udecimalPair struct {
	a udecimal.Decimal
	b udecimal.Decimal
}

type UdecimalTester struct {
	parseCases  []string
	stringCases []udecimal.Decimal
	addCases    []udecimalPair
	mulCases    []udecimalPair
	divCases    []udecimalPair
}

func NewUdecimalTester() Tester {
	return &UdecimalTester{}
}

func (self *UdecimalTester) Name() string {
	return "udecimal.Decimal"
}

func (self *UdecimalTester) Init() {
	self.parseCases = ParseCases

	self.stringCases = make([]udecimal.Decimal, len(ParseCases))
	for i, c := range ParseCases {
		d, err := udecimal.Parse(c)
		if err != nil {
			panic(err)
		}
		self.stringCases[i] = d
	}

	parsePairs := func(pairs []Pair) []udecimalPair {
		result := make([]udecimalPair, len(pairs))
		for i, p := range pairs {
			a, err := udecimal.Parse(p.A)
			if err != nil {
				panic(err)
			}
			b, err := udecimal.Parse(p.B)
			if err != nil {
				panic(err)
			}
			result[i] = udecimalPair{a, b}
		}
		return result
	}

	self.addCases = parsePairs(AddCases)
	self.mulCases = parsePairs(MulCases)
	self.divCases = parsePairs(DivCases)
}

func (self *UdecimalTester) RunParse() int64 {
	for range 100000 {
		for _, c := range self.parseCases {
			_, err := udecimal.Parse(c)
			if err != nil {
				panic(err)
			}
		}
	}
	return int64(len(self.parseCases) * 100000)
}

func (self *UdecimalTester) RunString() int64 {
	for range 100000 {
		for _, c := range self.stringCases {
			_ = c.String()
		}
	}
	return int64(len(self.stringCases) * 100000)
}

func (self *UdecimalTester) RunAdd() int64 {
	for range 100000 {
		for _, c := range self.addCases {
			_ = c.a.Add(c.b)
		}
	}
	return int64(len(self.addCases) * 100000)
}

func (self *UdecimalTester) RunMul() int64 {
	for range 100000 {
		for _, c := range self.mulCases {
			_ = c.a.Mul(c.b)
		}
	}
	return int64(len(self.mulCases) * 100000)
}

func (self *UdecimalTester) RunDiv() int64 {
	for range 100000 {
		for _, c := range self.divCases {
			_, err := c.a.Div(c.b)
			if err != nil {
				panic(err)
			}
		}
	}
	return int64(len(self.divCases) * 100000)
}
