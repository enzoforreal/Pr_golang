package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ComputeClosestToZero(ts []int) int {

	if len(ts) == 0 {
		return 0
	}

	closest := ts[0]
	for _, t := range ts {
		if abs(t) < abs(closest) || (abs(t) == abs(closest) && t > closest) {
			closest = t
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/* Ignore and do not change the code below */
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	ts := make([]int, n)
	for i := 0; i < n; i++ {
		ts[i], _ = strconv.Atoi(inputs[i])
	}
	stdoutWriter := os.Stdout
	os.Stdout = os.Stderr
	solution := ComputeClosestToZero(ts)
	os.Stdout = stdoutWriter
	fmt.Println(solution)
}

/* Ignore and do not change the code above */
