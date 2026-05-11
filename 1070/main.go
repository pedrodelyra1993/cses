/*
Print all evens in order, then all odds in order.

Within each half, consecutive elements differ by 2. The only risky point is the
junction: evens-first makes it (last_even, 1), guaranteeing a gap of at least 2
for n>=4. Odds-first would make it (last_odd, 2), which fails at n=4 since
last_odd=3 and |3-2|=1.
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

	if n == 2 || n == 3 {
		fmt.Fprintln(writer, "NO SOLUTION")
	} else {
		for even := 2; even <= n; even += 2 {
			fmt.Fprintf(writer, "%d ", even)
		}

		for odd := 1; odd <= n; odd += 2 {
			fmt.Fprintf(writer, "%d ", odd)
		}

		fmt.Fprintln(writer)
	}
}
