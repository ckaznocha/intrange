# intrange

intrange is a program for checking for loops that could use the [Go 1.22](https://go.dev/ref/spec#Go_1.22) integer
range feature.

## Installation

```bash
go install github.com/ckaznocha/intrange@latest
```

## Usage

```bash
intrange ./...
```

## Example

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```

Running `intrange` on the above code will produce the following output:

```bash
main.go:5:2: for loop can use an int range
```

The loop can be rewritten as:

```go
package main

import "fmt"

func main() {
    for i := range 10 {
        fmt.Println(i)
    }
}
```
