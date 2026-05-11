/*
The spiral can be viewed as concentric square layers.

For a position (x, y), the active layer is determined by max(x, y).
Each layer n has a diagonal anchor value:

	1, 3, 7, 13, 21, ...

The consecutive differences form an arithmetic sequence:

	2, 4, 6, 8, ...

For layer n, let k = n - 1. The k-th difference is:

	2 + 2(k - 1)

Summing the first k differences gives:

	S_k = k(a₁ + a_k)/2
	    = k(2 + 2k)/2
	    = k² + k

Thus the anchor of layer n is:

	1 + S_(n-1)
	= (n-1)² + (n-1) + 1

The final answer is obtained by moving an offset from this anchor.
The offset direction depends on parity because even and odd layers
spiral in opposite directions.

Coordinates are intentionally modeled as (x, y) internally for grid
intuition, even though the problem statement provides them as (y, x).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func getAnchor(n int64) int64 {
	k := n - 1
	return k*k + k + 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for range t {
		var x, y int64
		fmt.Fscan(reader, &x, &y)

		n := max(x, y)

		anchor := getAnchor(n)
		offset := n - min(x, y)
		if (n%2 == 0 && x < y) || (n%2 == 1 && x > y) {
			offset *= -1
		}

		result := anchor + offset
		fmt.Fprintln(writer, result)
	}
}
