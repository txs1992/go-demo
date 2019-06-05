package main

import (
	"fmt"
	"./myPkg"
)

func main() {
	fmt.Println("Hello, Go! " + myPkg.MyConst + " " + myPkg.MyType)
}