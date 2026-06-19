package main

import (
	"fmt"

	"github.com/yegorov/go-examples/internal/structs"

	"github.com/yegorov/go-examples/internal/lang"
)

func main() {
	fmt.Println("Run")
	lang.Case(10)

	stack := structs.BuildStack[string]()
	stack.Push("hello")
	stack.Push("world")
	stack.Print()
}
