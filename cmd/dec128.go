package main

import (
	"github.com/jokruger/dec128"
)

type dec128Pair struct {
	a dec128.Dec128
	b dec128.Dec128
}

type Dec128Tester struct {
	parseCases []string
	addCases   []dec128Pair
	mulCases   []dec128Pair
	divCases   []dec128Pair
}

func NewDec128Tester() Tester {
	return &Dec128Tester{}
}

func (self *Dec128Tester) Name() string {
	return "dec128.Dec128"
}

func (self *Dec128Tester) Init() {
	self.parseCases = ParseCases

	parsePairs := func(pairs []Pair) []dec128Pair {
		result := make([]dec128Pair, len(pairs))
		for i, p := range pairs {
			a := dec128.FromString(p.A)
			if a.IsNaN() {
				panic(a.ErrorDetails())
			}
			b := dec128.FromString(p.B)
			if b.IsNaN() {
				panic(b.ErrorDetails())
			}
			result[i] = dec128Pair{a, b}
		}
		return result
	}

	self.addCases = parsePairs(AddCases)
	self.mulCases = parsePairs(MulCases)
	self.divCases = parsePairs(DivCases)
}

func (self *Dec128Tester) RunParse() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.parseCases {
			d := dec128.FromString(c)
			if d.IsNaN() {
				panic(d.ErrorDetails())
			}
			n++
		}
	}
	return n
}

func (self *Dec128Tester) RunString() int64 {
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

func (self *Dec128Tester) RunAdd() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.addCases {
			_ = c.a.Add(c.b)
			n++
		}
	}
	return n
}

func (self *Dec128Tester) RunMul() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.mulCases {
			_ = c.a.Mul(c.b)
			n++
		}
	}
	return n
}

func (self *Dec128Tester) RunDiv() int64 {
	var n int64
	for range 100000 {
		for _, c := range self.divCases {
			_ = c.a.Div(c.b)
			n++
		}
	}
	return n
}
