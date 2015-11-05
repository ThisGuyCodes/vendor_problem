package main

import (
	"fmt"

	"a"
	"b"
	"c"
)

func main() {
	fmt.Println(c.Word)

	fmt.Println(a.F(c.Word))

	fmt.Println(b.F(c.Word))

	fmt.Println(b.F(a.F(c.Word)))
}
