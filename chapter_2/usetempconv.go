package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC).String())
	fmt.Println(tempconv.CToK(tempconv.BoilingC).String())
}
