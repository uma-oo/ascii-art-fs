package main

import (
	"log"
	"os"
	"fmt"
	function "ascii/functions"
)

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 1:
		fmt.Print(function.Fs(args[0]))
	case 2:
		fmt.Print(function.Fs(args[0], args[1]))
	default:
		log.Fatal("There is no arguments! Please provide a one!")

	}
}
