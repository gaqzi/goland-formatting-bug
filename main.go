package main

import (
	"fmt"
)

type builder struct{}

func (b builder) with(s string) builder {
	return b
}

func main() {
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	b := builder{}

	b.
		with("alalalalalalal"). // good explanatory comment
		with("bbababababababababbababab"). // less good explanatory comment
		with("cbcbcbbcbcbcbcbcbcbcbcbcbcbbccbcbc") // fantastic information

}
