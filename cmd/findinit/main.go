package main

import "fmt"
import "github.com/stephane-martin/go-findinit"

// Prints the type if the current system's init
func main() {
	fmt.Println(findinit.InitTypes[findinit.FindInit()])
}
