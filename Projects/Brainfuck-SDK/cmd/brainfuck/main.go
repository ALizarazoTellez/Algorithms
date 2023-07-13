package main

import (
	"log"
	"os"

	"github.com/ALizarazoTellez/Algorithms/Projects/Brainfuck-SDK/pkg/bf"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments!")
	}

	prog, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := bf.NewInterpreter(os.Stdout, os.Stdin).Execute(prog); err != nil {
		panic(err)
	}
}
