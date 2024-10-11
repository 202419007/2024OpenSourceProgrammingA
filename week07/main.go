package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	r := strings.NewReplacer("#", "o")
	fmt.Println(broken)
	fmt.Println(r.Replace(broken))
}
