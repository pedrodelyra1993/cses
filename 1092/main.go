/*
The core idea is that the sum of numbers from 1 to n can be described as:
1 + 2 + 3 + ... + n = n*(n+1)/2

If we want two sets of equal sum, we only need to find a subset of integers
that sums up to [n*(n+1)]/4.

This is the classic coin change problem. Since the numbers from 1 to n form a canonical coin system
that can be solved using the greedy approach, there is no need to the DP solution.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int64
	fmt.Fscan(reader, &n)
	gauss := n * (n + 1) / 2
	if gauss%2 != 0 {
		fmt.Fprintln(writer, "NO")
		return
	}

	seen := make([]bool, n+1)
	for i, current := n, gauss/2; i >= 1; i-- {
		if i > current {
			continue
		}

		current -= i
		seen[i] = true
	}

	fmt.Fprintln(writer, "YES")

	s1 := []int64{}
	s2 := []int64{}
	for i := int64(1); i <= n; i++ {
		if seen[i] {
			s1 = append(s1, i)
		} else {
			s2 = append(s2, i)
		}
	}

	fmt.Fprintln(writer, len(s1))
	for _, n := range s1 {
		fmt.Fprintf(writer, "%d ", n)
	}
	fmt.Fprintln(writer)

	fmt.Fprintln(writer, len(s2))
	for _, n := range s2 {
		fmt.Fprintf(writer, "%d ", n)
	}
	fmt.Fprintln(writer)

	sum := func(arr []int64) int64 {
		s := int64(0)
		for _, n := range arr {
			s += n
		}
		return s
	}

	if sum(s1) != sum(s2) {
		panic("FAILED!")
	}
}
