package staircaseproblem

import "fmt"

func StaircaseProblem(n int32) {
	var i, j int32
	for i = 1; i <= n; i++ {
		for j = n; j >= 1; j-- {
			if i >= j {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
