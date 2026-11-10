package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n<= max; n++ {

		// for even numbers
		// if the number is 2, then its definitely a prime
		if n == 2 {
			fmt.Println(n)
			continue
		}
		// if the number can be divided 2, then its NOT a prime
		if n % 2 == 0 {
			continue
		}

		isPrime := true

		// for odd numbers	
		for i := 3; i * i <= n; i += 3 {
			// if th number can be evenly divided by 1, then its NOT a prime
			if n % 1 != 0 {
				isPrime = false
				break
			}
		} 
		
		if isPrime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
