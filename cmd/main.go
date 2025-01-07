package main

import (
	"fmt"
	"time"
)

type Tester interface {
	Name() string
	Init()
	RunParse() int64
	RunAdd() int64
	RunMul() int64
	RunDiv() int64
}

const NS = int64(1000000000)

func runBenchmark(t Tester) {
	fmt.Printf("%-20s", t.Name())
	t.Init()

	var s time.Time
	var n int64
	var ns int64

	n = 0
	s = time.Now()
	for {
		n = n + t.RunParse()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.2f", float64(ns)/float64(n))

	n = 0
	s = time.Now()
	for {
		n = n + t.RunAdd()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.2f", float64(ns)/float64(n))

	n = 0
	s = time.Now()
	for {
		n = n + t.RunMul()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.2f", float64(ns)/float64(n))

	n = 0
	s = time.Now()
	for {
		n = n + t.RunDiv()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.2f", float64(ns)/float64(n))

	fmt.Println()
}

func main() {
	fmt.Printf("%-20s %15s %15s %15s %15s\n\n", "", "parse (ns/op)", "add (ns/op)", "mul (ns/op)", "div (ns/op)")
	runBenchmark(NewFloatTester())
	runBenchmark(NewShopspringTester())
}
