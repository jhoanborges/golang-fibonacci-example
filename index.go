package main

import (
	"fmt"
	"math/big"
)

func main()  {
	cont := 0
	var limit int
	limit = 500

	for i := 1; i <= limit; i++ {
		fmt.Println("INDEX: ",i, "Fibonacci number: ",FibonacciBig( uint(i)))
		cont += i
	}
	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() // wait for Enter Key

}

// FibonacciBig calculates Fibonacci number using bit.Int.
// For the sequence numbers below 94, it is recommended to use Fibonacci function as it is more efficient.
func FibonacciBig(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}
	return n1
}
//i used this package as a reference https://github.com/t-pwk/go-fibonacci
//https://blog.abelotech.com/posts/first-500-fibonacci-numbers/ first  500 fibonacci numbers