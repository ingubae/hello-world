package main

import (
	"fmt"

	"github.com/ingubae/stringutil"
)

func main() {
	fmt.Println("Hello World")
	val := stringutil.Reverse("ABCDEF")
	fmt.Println(val)
}
