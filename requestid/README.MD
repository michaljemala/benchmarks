go test -bench BenchmarkUUID -benchmem -run ^$
```
BenchmarkUUIDUlid-8     	10000000	       144 ns/op	      48 B/op	       2 allocs/op
BenchmarkUUIDGoogle-8   	10000000	       192 ns/op	      48 B/op	       1 allocs/op
BenchmarkUUIDSatori-8   	10000000	       156 ns/op	      48 B/op	       1 allocs/op
BenchmarkUUIDAtomic-8   	20000000	        89.0 ns/op	     128 B/op	       2 allocs/op
PASS
ok  	github.com/michaljemala/benchmarks/requestid	7.329s
```
