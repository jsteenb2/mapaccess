### Run the benchmarks

1. go test -bench=. -cpuprofile=cpu.out
2. go test -bench=. -memprofile=mem.out

### Fire up pprof to get the skinny

1. go tool pprof mapaccess.test cpu.out
2. go tool pprof --alloc_space mapaccess.test mem0.out

### Export cpu profile web view to pdf

1. go tool pprof --pdf mapaccess.test cpu.out > cpu.pdf
