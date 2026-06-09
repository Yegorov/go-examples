## Go examples

### How to

```
git clone git@github.com:Yegorov/go-examples.git
cd go-examples

go build -o build/hello cmd/hello.go
./build/hello

go run cmd/share/<file>.go
```

### Run integration tests
```
go test -tags=integration -v ./cmd/...
```

### Run benchmarks
```
go test -tags=integration -v ./cmd/... -bench=.
go test -tags=integration -v ./cmd/... -bench=. -benchmem
```

```
go install golang.org/x/perf/cmd/benchstat@latest
go test -tags=integration -v ./cmd/... -bench=. -count=10 > old.txt
go test -tags=integration -v ./cmd/... -bench=. -count=10 > new.txt
benchstat old.txt new.txt
rm old.txt new.txt
```
