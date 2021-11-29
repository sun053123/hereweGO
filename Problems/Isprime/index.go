package main

import "fmt"

func main() {
	n := 50
	a := isPrime(n)
	fmt.Println(a)
}

func isPrime(n int) bool {

	var prime []int
	if n > 1 {
		prime = append(prime, 1)
	}
	if n == 1 {
		return false
	}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			prime = append(prime, i)
		}
	}

	if len(prime) >= 3 {
		return false
	} else {
		return true
	}
}
