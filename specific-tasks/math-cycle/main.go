package main

import "fmt"

const transfer = 23

func getSlicePart(infra []int, i int) []int {
	lenPart := i + transfer
	if i+transfer >= len(infra) {
		lenPart = len(infra)
	}
	return infra[i:lenPart]
}

func main() {

	lenOb := 137

	ob := make([]int, lenOb)

	for i := 0; i < lenOb; i++ {
		ob[i] = i
	}

	// fmt.Println(ob)

	for i := 0; i < lenOb; i += transfer {
		fmt.Println(getSlicePart(ob, i))
	}
}
