package main

import (
	"finautils"
	"fmt"
)

func main() {
	intRef, err := finautils.CreateIntRef("1009")
	if err != nil {
		panic(err.Error())
	}
	if !finautils.ValidateFinnishReference("10091") {
		fmt.Println("1009 is not valid")
	}
	fmt.Println("INT REF:", intRef)
}
