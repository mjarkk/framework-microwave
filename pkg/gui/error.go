package gui

import (
	"fmt"
	"os"
)

// CritErr logs out an error and kills the program
func CritErr(erroMsg string) {
	fmt.Println("Critical error:")
	fmt.Println(erroMsg)
	os.Exit(1)
}
