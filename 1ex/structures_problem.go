package main

import "fmt"

type Vehicle struct {
	Name  string
	Price int
}

func (v *Vehicle) Drive(speed int) {
	fmt.Println(speed > 0)
}

type Car struct {
	Vehicle
}

func main() {
	merc := Car{}
	merc.Name = "Mercedes-Benz"
	merc.Price = 100000000
	merc.Drive(10)
}
