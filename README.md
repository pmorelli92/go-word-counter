## Go Word Counter Sync Or Concurrent

Read more in my [blog](https://devandchill.com/posts/2019/04/word-counter-sync-or-concurrent/) 😁


# Run

    go test -bench=. --benchmem

    goos: darwin
    goarch: amd64
    pkg: github.com/pmorelli92/go-word-counter
    BenchmarkWordCounter-8             	     100	  17365420 ns/op	 2702199 B/op	  129059 allocs/op
    BenchmarkWordCounterConcurrent-8   	      30	  52831085 ns/op	 7004061 B/op	  137706 allocs/op
    BenchmarkWordCounterStreams-8      	     300	   5398546 ns/op	  426424 B/op	   50175 allocs/op
    PASS
    ok  	github.com/pmorelli92/go-word-counter	5.593s
