package main

import "fmt"

func main() {
	//  inside scope of (x)
	x := "This is variable x"
	fmt.Printf("%s inside scope of x.\n", x)
	{
		//  inside scope of (y)
		fmt.Printf("%s inside scope of y.\n", x)
		var y = "I have no idea what I'm doing :)"
		fmt.Printf("%s, but this line is inside scope of y.", y)
	}

}
