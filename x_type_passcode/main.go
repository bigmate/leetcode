package main

import "fmt"

func calcSteps(passcode, numpad string) int {
	steps := 0
	pad := parseNumpad(numpad)
	adj := adjacencyMap(pad)

	for i := 1; i < len(passcode); i++ {
		pre := passcode[i-1] - '0'
		cur := passcode[i] - '0'

		if pre == cur {
			continue
		}

		_, ok := adj[pre][cur]

		if ok {
			steps++
		} else {
			steps += 2
		}

	}

	return steps
}

func adjacencyMap(pad [3][3]uint8) [10]map[uint8]struct{} {
	adj := [10]map[uint8]struct{}{}

	for i := 1; i < 10; i++ {
		adj[i] = make(map[uint8]struct{}, 4)
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			n := pad[y][x]
			if x > 0 {
				adj[n][pad[y][x-1]] = struct{}{} // left
			}
			if x < 2 {
				adj[n][pad[y][x+1]] = struct{}{} // right
			}
			if y > 0 {
				adj[n][pad[y-1][x]] = struct{}{} // top
			}
			if y < 2 {
				adj[n][pad[y+1][x]] = struct{}{} // bottom
			}
			if x > 0 && y > 0 {
				adj[n][pad[y-1][x-1]] = struct{}{} // top left
			}
			if x < 2 && y < 2 {
				adj[n][pad[y+1][x+1]] = struct{}{} // bottom right
			}
			if x < 2 && y > 0 {
				adj[n][pad[y-1][x+1]] = struct{}{} // top right
			}
			if x > 0 && y < 2 {
				adj[n][pad[y+1][x-1]] = struct{}{} // bottom left
			}
		}
	}

	return adj
}

func parseNumpad(numpad string) [3][3]uint8 {
	pad := [3][3]uint8{}

	for i := 0; i < 9; i++ {
		pad[i/3][i%3] = numpad[i] - '0'
	}

	return pad
}

func main() {
	fmt.Println(calcSteps("8978234", "672458139"))
}

// 8978234
//
// 6 7 2
// 4 5 8
// 1 3 9
//
// Ans: 8
