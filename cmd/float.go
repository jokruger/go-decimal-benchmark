package main

import (
	"strconv"
)

type floatPair struct {
	a float64
	b float64
}

type FloatTester struct {
	parseCases  []string
	stringCases []float64
	addCases    []floatPair
	mulCases    []floatPair
	divCases    []floatPair
}

func NewFloatTester() Tester {
	return &FloatTester{}
}

func (self *FloatTester) Name() string {
	return "float64 (baseline)"
}

func (self *FloatTester) Init() {
	self.parseCases = ParseCases

	self.stringCases = make([]float64, len(ParseCases))
	for i, c := range ParseCases {
		d, err := strconv.ParseFloat(c, 64)
		if err != nil {
			panic(err)
		}
		self.stringCases[i] = d
	}

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
	for range 100000 {
		for _, c := range self.stringCases {
			_ = strconv.FormatFloat(c, 'f', -1, 64)
		}
	}
	return int64(len(self.stringCases) * 100000)
}

func (self *FloatTester) RunAdd() int64 {
	for range 100000 {
		for _, c := range self.addCases {
			_ = c.a + c.b
		}
	}
	return int64(len(self.addCases) * 100000)
}

func (self *FloatTester) RunMul() int64 {
	for range 100000 {
		for _, c := range self.mulCases {
			_ = c.a * c.b
		}
	}
	return int64(len(self.mulCases) * 100000)
}

func (self *FloatTester) RunDiv() int64 {
	for range 100000 {
		for _, c := range self.divCases {
			_ = c.a / c.b
		}
	}
	return int64(len(self.divCases) * 100000)
}
