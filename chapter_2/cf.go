package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf:%v\n", err)
				os.Exit(1)
			}
			ShowResult(t)
		}
	} else {
		for _, arg := range args {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf:%v\n", err)
				os.Exit(1)
			}
			ShowResult(t)
		}
	}

	// for _, arg := range os.Args[1:] {
	// 	t, err := strconv.ParseFloat(arg, 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "cf:%v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	f := tempconv.Fahrenheit(t)
	// 	c := tempconv.Celsius(t)

	// 	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	// }
}

func ShowResult(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
