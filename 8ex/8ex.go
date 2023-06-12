package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var z int64 = 2012
	var i int = 9

	new, err := reverseVar(z, i)
	if err != nil {
		log.Fatal("error")
	}

	fmt.Printf("%d -> %b\n", z, z)
	fmt.Printf("%b -> %d\n", new, new)
}

func reverseVar(z int64, i int) (int64, error) {
	var new int64
	var tmp int64

	if i < 1 || i > 63 {
		return 0, errors.New("oh shit")
	}

	if z < 0 {
		tmp = -z
	} else {
		tmp = z
	}

	i--
	switch tmp & (1 << i) {
	case 0:
		new = tmp | (1 << i)
	default:
		new = tmp & ^(1 << i)
	}

	if z < 0 {
		return -new, nil
	}
	return new, nil
}
