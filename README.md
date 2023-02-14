# golang-prime

This package provides some functions for working with prime numbers.

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