package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("\n")
	fmt.Printf("This \t is \t a \t space \t test! \n \n")
	fmt.Printf("Derp derp line 2 \n \n")
	fmt.Printf("Derp derp derp line 3 \n \n")
	fmt.Printf("Derp derp derp derp line 4 \n \n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Launch Missile? [y/n]:")
	_, _ = reader.ReadString('\n')
	fmt.Println("FUCK YOU")
}
