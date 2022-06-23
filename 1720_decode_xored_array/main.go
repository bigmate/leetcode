package main

import "fmt"

func decode(encoded []int, first int) []int {
	res := make([]int, 0, len(encoded)+1)
	res = append(res, first)

	for _, code := range encoded {
		res = append(res, code^res[len(res)-1])
	}

	return res
}

func main() {
	fmt.Println(decode([]int{}, 0))
}
