package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var pc [256]byte

func main() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		v, _ := strconv.ParseUint(input.Text(), 0, 64)

		start := time.Now()
		fmt.Printf("%d = %d\n", v, PopCount(v))
		fmt.Printf("%dms elapsed\n", time.Since(start).Milliseconds())

		start = time.Now()
		fmt.Printf("%d = %d\n", v, PopCount2(v))
		fmt.Printf("%dms elapsed\n", time.Since(start).Milliseconds())
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	c := 0
	for ; x > 0; x >>= 1 {
		if x&1 == 1 {
			c++
		}
	}
	return c
}
