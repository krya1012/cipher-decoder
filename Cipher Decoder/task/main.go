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

func caesar(text string, shift int) string {
	shift = ((shift % 26) + 26) % 26
	out := make([]byte, len(text))
	for i, ch := range text {
		switch {
		case ch >= 'a' && ch <= 'z':
			out[i] = byte((int(ch-'a')+shift)%26 + 'a')
		case ch >= 'A' && ch <= 'Z':
			out[i] = byte((int(ch-'A')+shift)%26 + 'A')
		default:
			out[i] = byte(ch)
		}
	}
	return string(out)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var g, p int
	_, _ = fmt.Sscanf(scanner.Text(), "g is %d and p is %d", &g, &p)
	fmt.Println("OK")

	b := rand.Intn(p-2) + 2 // 1 < b < p
	B := modExp(g, b, p)

	scanner.Scan()
	var A int
	_, _ = fmt.Sscanf(scanner.Text(), "A is %d", &A)

	S := modExp(A, b, p)

	fmt.Printf("B is %d\n", B)

	shift := S % 26
	fmt.Println(caesar("Will you marry me?", shift))

	scanner.Scan()
	reply := caesar(scanner.Text(), -shift)

	switch reply {
	case "Yeah, okay!":
		fmt.Println(caesar("Great!", shift))
	case "Let's be friends.":
		fmt.Println(caesar("What a pity!", shift))
	}
}
