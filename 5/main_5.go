package main

import "fmt"

func sameElements(a, b []int) (bool, []int) {
	counts := make(map[int]int)
	for _, v := range a {
		counts[v]++
	}

	result := []int{}
	for _, v := range b {
		if counts[v] > 0 {
			result = append(result, v)
			counts[v]--
		}
	}

	if len(result) == 0 {
		return false, nil
	}
	return true, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(sameElements(a, b))
}
