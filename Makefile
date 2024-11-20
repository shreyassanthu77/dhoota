mode=debug
target=x86_64-linux-musl
export CC=zig cc -target $(target)
export CXX=zig c++ -target $(target)
export CGO_ENABLED=1

tags='static'
extldflags="-static"
ifeq ($(mode), release)
	extldflags="-static -O3 -s"
endif
ldflags='-linkmode=external -extldflags $(extldflags)'

src=$(wildcard *.go)
build: $(src)
	go build -ldflags $(ldflags) -tags $(tags) -o bin/main .

run: build
	./bin/main

watch:
	gow -c run -ldflags $(ldflags) -tags $(tags) .
