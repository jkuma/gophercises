package main

import (
	"fmt"
)

type Animal interface {
	Hello() string
}

type Cat struct {
	Color string
}

func (c Cat) Hello() string {
	return "Miaaaouuuuu"
}

type MecanicCat struct {
	Cat
}

func catHello(a Animal) {
	fmt.Println(a.Hello())
}

func main() {
	cat := Cat{Color: "black"}
	catHello(cat)

	mc := MecanicCat{Cat: cat}
	catHello(mc)
}
