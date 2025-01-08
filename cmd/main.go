package main

import (
	"fmt"
	"time"

	. "github.com/klauspost/cpuid/v2"
)

type Tester interface {
	Name() string
	Init()
	RunString() int64
	RunParse() int64
	RunAdd() int64
	RunMul() int64
	RunDiv() int64
}

const NS = int64(1000000000)

func runBenchmark(t Tester) {
	fmt.Printf("%-30s", t.Name())
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
	fmt.Printf(" %15.3f", float64(ns)/float64(n))

	n = 0
	s = time.Now()
	for {
		n = n + t.RunString()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.3f", float64(ns)/float64(n))

	n = 0
	s = time.Now()
	for {
		n = n + t.RunAdd()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.3f", float64(ns)/float64(n))

	n = 0
	s = time.Now()
	for {
		n = n + t.RunMul()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.3f", float64(ns)/float64(n))

	n = 0
	s = time.Now()
	for {
		n = n + t.RunDiv()
		ns = time.Since(s).Nanoseconds()
		if ns >= NS {
			break
		}
	}
	fmt.Printf(" %15.3f", float64(ns)/float64(n))

	fmt.Println()
}

func main() {
	fmt.Println("CPU:", CPU.BrandName)
	fmt.Println()
	fmt.Printf("%-30s %15s %15s %15s %15s %15s\n\n", "", "parse (ns/op)", "string (ns/op)", "add (ns/op)", "mul (ns/op)", "div (ns/op)")
	runBenchmark(NewFloatTester())
	runBenchmark(NewDec128Tester())
	runBenchmark(NewUdecimalTester())
	runBenchmark(NewAlpacadecimalTester())
	runBenchmark(NewShopspringTester())
}
