package main

import "fmt"

func resultSlice(slice1 []string, slice2 []string) []string {
	resultSlice := []string{}
	exist := make(map[string]bool)
	for _, v := range slice2 {
		exist[v] = true
	}
	for _, v := range slice1 {
		if !exist[v] {
			resultSlice = append(resultSlice, v)
		}
	}
	return resultSlice
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(resultSlice(slice1, slice2))
}
