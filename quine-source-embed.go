package main

import (
	_ "embed"
	"fmt"
)

//go:embed quine-source-embed.go
var program string

func main() {
	fmt.Print(program)
}
