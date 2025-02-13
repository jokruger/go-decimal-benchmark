# go-decimal-benchmark

Benchmarks for Go decimal packages

## How to run

```bash
go run ./cmd
```

## Results

```
CPU: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

                                 parse (ns/op)  string (ns/op)     add (ns/op)     mul (ns/op)     div (ns/op)

float64 (baseline)                      30.303          82.617           0.352           0.362           0.344
dec128.Dec128                           13.314          33.699          11.586           7.581          34.058
udecimal.Decimal                        23.258          41.481          12.151          11.210          40.126
alpacadecimal.Decimal                   88.490          77.343         205.212          60.344         454.116
shopspring.Decimal                     151.691         171.424         216.789          66.622         435.903
```
