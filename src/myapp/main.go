package main

import (
	"fmt"
	"mylib"
	"mylib2"
)

func main() {
	fmt.Println("mylib.Foo(): ", mylib.Foo())
	fmt.Println("mylib.Bar(): ", mylib2.Bar())
}
