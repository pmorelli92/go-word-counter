## Go Word Counter Sync Or Concurrent

Read more in my [blog](https://devandchill.com/posts/2019/04/word-counter-sync-or-concurrent/) 😁


# Run

```
goos: darwin
goarch: amd64
PASS
benchmark                           iter     time/iter    bytes alloc             allocs
---------                           ----     ---------    -----------             ------
BenchmarkWordCounter-12               45   25.88 ms/op   4354027 B/op   124918 allocs/op
BenchmarkWordCounterConcurrent-12     19   59.70 ms/op   5711640 B/op   138290 allocs/op
BenchmarkWordCounterStreams-12       198    5.93 ms/op    330114 B/op    44844 allocs/op
ok      _/Users/pabmor/some/go-word-counter     5.952s
```
