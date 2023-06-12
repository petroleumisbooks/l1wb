package main

import (
	"errors"
	"fmt"
)

func main() {
	var a interface{}

	a = 1

	example, _ := define(a)

	fmt.Println(example)
}

func define(a interface{}) (string, error) {
	var Type string
	switch a.(type) {
	case int:
		Type = "int"
	case string:
		Type = "string"
	case bool:
		Type = "bool"
	default:
		return "", errors.New("hz")
	}
	return Type, nil
}
