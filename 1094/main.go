/*
Scan left to right, comparing each element to its predecessor.

When an element is smaller, raise it to match and accumulate the difference
as the move cost.
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

	var n int
	fmt.Fscan(reader, &n)

	x := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	moves := int64(0)
	for i := 1; i < n; i++ {
		if x[i] < x[i-1] {
			moves += x[i-1] - x[i]
			x[i] = x[i-1]
		}
	}

	fmt.Fprintln(writer, moves)
}
