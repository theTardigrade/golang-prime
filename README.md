# golang-prime

This package provides some functions for working with prime numbers.

[![Go Reference](https://pkg.go.dev/badge/github.com/theTardigrade/golang-prime.svg)](https://pkg.go.dev/github.com/theTardigrade/golang-prime) [![Go Report Card](https://goreportcard.com/badge/github.com/theTardigrade/golang-prime)](https://goreportcard.com/report/github.com/theTardigrade/golang-prime)

## Example

```golang
package main

import (
	"fmt"

	prime "github.com/theTardigrade/golang-prime"
)

func main() {
	magicNumber := 3

	fmt.Println(prime.Is(magicNumber))   // true
	fmt.Println(prime.Next(magicNumber)) // 5

	fmt.Println("*****")

	magicNumber = 120

	fmt.Println(prime.Is(magicNumber))   // false
	fmt.Println(prime.Next(magicNumber)) // 127
}
```

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)