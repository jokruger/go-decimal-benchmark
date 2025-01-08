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

float64 (baseline)                      33.408          92.405           0.335           0.340           0.331
dec128.Dec128                           13.986          36.404          10.518           7.637          34.129
udecimal.Decimal                        22.383          44.740          11.998          11.141          40.701
alpacadecimal.Decimal                   90.959          83.291         222.275          70.552         481.113
shopspring.Decimal                     160.160         183.984         241.129          74.726         451.901
```
