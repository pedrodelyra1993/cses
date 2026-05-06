/*
The Collatz conjecture, also known as the 3n+1 conjecture, is a famous unsolved problem in mathematics.
It states that for any positive integer n, repeatedly applying the following operations will eventually
reach 1: if n is even, divide it by 2; if n is odd, multiply it by 3 and add 1. Although the conjecture
has been verified for very large numbers, it remains unproven for all integers.

This problem requires simulating the Collatz sequence for a given n, printing each step until reaching 1.
It demonstrates the conjecture's process without needing to prove it.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func nextCollatz(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int64
	fmt.Fscan(reader, &n)
	fmt.Fprint(writer, n)
	for n > 1 {
		n = nextCollatz(n)
		fmt.Fprint(writer, " ", n)
	}
}
