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

float64 (baseline)                      30.897          84.905           0.333           0.342           0.328
dec128.Dec128                           14.024          33.683           9.975           6.569          35.116
udecimal.Decimal                        23.302          41.854          12.226          11.346          40.877
alpacadecimal.Decimal                   89.528          78.884         206.393          60.364         451.828
shopspring.Decimal                     152.263         169.300         218.909          65.241         428.002
```
