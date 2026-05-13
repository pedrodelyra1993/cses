package main

import "fmt"

const mod int64 = 1e9 + 7

func multmod(a int64, b int64) int64 {
	return ((a % mod) * (b % mod)) % mod
}

func binexp(a int64, b int64) int64 {
	if b == 0 {
		return 1
	}

	k := binexp(a, b/2)
	m := multmod(k, k)
	if b%2 == 0 {
		return m
	}

	return multmod(a, m)
}

func main() {
	var n int64
	fmt.Scan(&n)

	result := binexp(2, n)
	fmt.Println(result)
}
