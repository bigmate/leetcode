package main

import (
	"fmt"
)

func numSteps(s string) (steps int) {
	bits := []byte(s)
	i := len(bits) - 1
	for i > 0 {
		if bits[i] == '0' {
			steps++
			i--
		} else {
			steps++
			j := i
			for j > 0 {
				if bits[j] == '0' {
					bits[j] = '1'
					break
				}
				j--
			}
			if j <= 0 {
				steps++
			}
			steps += i - j
			i = j
		}
	}
	return
}

func main() {
	steps := numSteps("10011101")
	fmt.Println(steps)
}

/*  Example:
0) 10011101
1) 10011110
2) 1001111
3) 1010000
4) 101000
5) 10100
6) 1010
7) 101
8) 110
9) 11
10) 100
11) 10
12) 1
*/
