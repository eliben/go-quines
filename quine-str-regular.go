package main

import "fmt"

func main() {
	t := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tt := %q\n\tfmt.Printf(t, t)\n}\n"
	fmt.Printf(t, t)
}
