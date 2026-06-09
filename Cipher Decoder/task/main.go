package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func modExp(base, exp, mod int) int {
	c := 1
	for i := 0; i < exp; i++ {
		c = (c * base) % mod
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var g, p int
	fmt.Sscanf(scanner.Text(), "g is %d and p is %d", &g, &p)
	fmt.Println("OK")

	b := rand.Intn(p-2) + 2 // 1 < b < p
	B := modExp(g, b, p)

	scanner.Scan()
	var A int
	fmt.Sscanf(scanner.Text(), "A is %d", &A)

	S := modExp(A, b, p)

	fmt.Printf("B is %d\n", B)
	fmt.Printf("A is %d\n", A)
	fmt.Printf("S is %d\n", S)
}
