package main

import (
	"fmt"

	"./anotherpkgvar" //don't forget to capitalize for visibility
)

func main() {

	fmt.Printf("%s, %s!\n\nHow are you doing today?", anotherpkgvar.B, anotherpkgvar.A)

}
