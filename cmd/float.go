package main

import (
	"strconv"
)

type floatPair struct {
	a float64
	b float64
}

type FloatTester struct {
	parseCases []string
	addCases   []floatPair
	mulCases   []floatPair
	divCases   []floatPair
}

func NewFloatTester() Tester {
	return &FloatTester{}
}

func (self *FloatTester) Name() string {
	return "float64 (baseline)"
}

func (self *FloatTester) Init() {
	self.parseCases = ParseCases

	parsePairs := func(pairs []Pair) []floatPair {
		result := make([]floatPair, len(pairs))
		for i, p := range pairs {
			a, err := strconv.ParseFloat(p.A, 64)
			if err != nil {
				panic(err)
			}
			b, err := strconv.ParseFloat(p.B, 64)
			if err != nil {
				panic(err)
			}
			result[i] = floatPair{a, b}
		}
		return result
	}

	self.addCases = parsePairs(AddCases)
	self.mulCases = parsePairs(MulCases)
	self.divCases = parsePairs(DivCases)
}

func (self *FloatTester) RunParse() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.parseCases {
			_, err := strconv.ParseFloat(c, 64)
			if err != nil {
				panic(err)
			}
			n++
		}
	}
	return n
}

func (self *FloatTester) RunString() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.addCases {
			_ = strconv.FormatFloat(c.a, 'f', -1, 64)
			_ = strconv.FormatFloat(c.b, 'f', -1, 64)
			n += 2
		}
	}
	return n
}

func (self *FloatTester) RunAdd() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.addCases {
			_ = c.a + c.b
			n++
		}
	}
	return n
}

func (self *FloatTester) RunMul() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.mulCases {
			_ = c.a * c.b
			n++
		}
	}
	return n
}

func (self *FloatTester) RunDiv() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.divCases {
			_ = c.a / c.b
			n++
		}
	}
	return n
}
