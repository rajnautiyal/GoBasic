package main

import (
	"fmt"
	"log"

	"github.com/rajnautiyal/myniceprogram/helpers"
)

func main() {
	fmt.Println("hello raj")
	log.Println("this is the programming for the oacjergs")

	var myVar helpers.SomeType
	myVar.Name = "rajendra"
	myVar.Age = "12"

	log.Println(myVar.Age, " now the name is ", myVar.Name)
	helpers.MultipleDiminesionSlice()
}
