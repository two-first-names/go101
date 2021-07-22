// Usage: go run cmd/fizzbuzz/fizzbuzz.go --max 10
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main() {
	max := flag.Int("max", 100, "Maximum number to fizzbuzz to")
	flag.Parse()

	for i := 1; i <= *max; i++ {
		fb, err := fizzbuzz(i)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(i, fb)
	}
}

func fizzbuzz(i int) (string, error) {
	fizz := "fizz"
	buzz := "buzz"

	if i < 0 {
		return "", errors.New("cannot do fizzbuzz on negative numbers")
	}

	if i % 3 == 0 && i % 5 == 0 {
		return fizz+buzz, nil
	} else if i % 3 == 0 {
		return fizz, nil
	} else if i % 5 == 0 {
		return buzz, nil
	} else {
		return "", nil
	}
}