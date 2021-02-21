package main

import (
	"fmt"
	"log"
)

const N = 1000
const M = 1000

var x = []int{0, 0, 1, -1}
var y = []int{1, -1, 0, 0}

var count [N][M]int
var used [N][M]bool
var a [N][M]rune
var ans = 0

func dfs(row int, col int, n int, m int, c int) {

	c += 1
	if c > (n + m - 1) {
		return
	}

	if row == 0 && col == m-1 {
		ans++
		return
	}
	if used[row][col] {
		return
	}

	used[row][col] = true
	for i := 0; i < 4; i++ {
		if (row+x[i] >= 0 && row+x[i] < n) && (col+y[i] >= 0 && col+y[i] < m) && (a[row+x[i]][col+y[i]] == '1') {
			dfs(row+x[i], col+y[i], n, m, c)
		}
	}
	used[row][col] = false
	return

}

func main() {
	var n int
	var m int

	_, err := fmt.Scan(&n, &m)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, err := fmt.Scan(&a[i][j])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	dfs(n-1, 0, n, m, 0)
	if ans == 0 {
		print("impossible")
	} else {
		print(ans)
	}

}
