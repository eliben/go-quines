package main

import "fmt"

var data = []rune{125, 10, 10, 102, 117, 110, 99, 32, 109, 97, 105, 110, 40, 41, 32, 123, 10, 9, 102, 109, 116, 46, 80, 114, 105, 110, 116, 102, 40, 34, 112, 97, 99, 107, 97, 103, 101, 32, 109, 97, 105, 110, 92, 110, 92, 110, 105, 109, 112, 111, 114, 116, 32, 92, 34, 102, 109, 116, 92, 34, 92, 110, 92, 110, 118, 97, 114, 32, 100, 97, 116, 97, 32, 61, 32, 91, 93, 114, 117, 110, 101, 123, 34, 41, 10, 10, 9, 102, 111, 114, 32, 105, 44, 32, 100, 32, 58, 61, 32, 114, 97, 110, 103, 101, 32, 100, 97, 116, 97, 32, 123, 10, 9, 9, 102, 109, 116, 46, 80, 114, 105, 110, 116, 102, 40, 34, 37, 100, 34, 44, 32, 100, 41, 10, 9, 9, 105, 102, 32, 105, 32, 60, 32, 108, 101, 110, 40, 100, 97, 116, 97, 41, 45, 49, 32, 123, 10, 9, 9, 9, 102, 109, 116, 46, 80, 114, 105, 110, 116, 102, 40, 34, 44, 32, 34, 41, 10, 9, 9, 125, 10, 9, 125, 10, 10, 9, 102, 111, 114, 32, 95, 44, 32, 100, 32, 58, 61, 32, 114, 97, 110, 103, 101, 32, 100, 97, 116, 97, 32, 123, 10, 9, 9, 102, 109, 116, 46, 80, 114, 105, 110, 116, 102, 40, 34, 37, 99, 34, 44, 32, 100, 41, 10, 9, 125, 10, 125, 10}

func main() {
	fmt.Printf("package main\n\nimport \"fmt\"\n\nvar data = []rune{")

	for i, d := range data {
		fmt.Printf("%d", d)
		if i < len(data)-1 {
			fmt.Printf(", ")
		}
	}

	for _, d := range data {
		fmt.Printf("%c", d)
	}
}