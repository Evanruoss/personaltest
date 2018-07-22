package main

import (
	"fmt"

	"./anotherpkgvar"
)

func main() {

	fmt.Printf("%s, %s", anotherpkgvar.b, anotherpkgvar.a)

}
