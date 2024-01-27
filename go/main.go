package main

import (
	"fmt"
)

type Person struct {
	name       string
	age        int
	profession string
	isMAle     bool
}

func main() {
	primitvo := Person{name: "John", age: 20, profession: "Programmer", isMAle: true}
	fmt.Println(primitvo)

}
