# Go Atomic

Atomic is a Go library that implements atomic types (eg AtomicCounter, AtomicInt, AtomicBool).

## Getting Started

Just a quick example how to use the library:

#### main.go
```
package main

import (
	"fmt"

	"github.com/jattschneider/atomic"
)

func main() {
	c := &atomic.AtomicCounter{}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			c.Incr()
			wg.Done()
		}()
	}
	wg.Wait()

	c.Incr()
}

```

```
$ go run main.go
```

### Installing

```
go get -v github.com/jattschneider/atomic
```

## Built With

* [Go](https://golang.org/) - The Go Programming Language

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jattschneider/argonauts/tags). 

## Authors

* **José Augusto Schneider** - *Initial work* - [jattschneider](https://github.com/jattschneider)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
