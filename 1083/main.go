/*
This problem can be solved using the formula for the sum of the first n natural numbers.

The expected sum of numbers from 1 to n is n(n+1)/2. Since exactly one number is missing
from the input, we can compute the sum of the given numbers and subtract it from the expected
sum to find the missing value.

This approach runs in O(n) time and uses O(1) extra space, avoiding the need for additional
data structures like maps or arrays.
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

	var sum int64
	for i := 0; i < n-1; i++ {
		var num int
		fmt.Fscan(reader, &num)
		sum += int64(num)
	}

	gauss := int64(n) * int64(n+1) / 2
	result := gauss - sum

	fmt.Fprintln(writer, result)
}
