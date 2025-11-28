package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(originalSlice []int) []int {
	var newSlice []int
	for i := 0; i < len(originalSlice); i++ {
		if originalSlice[i]%2 == 0 {
			newSlice = append(newSlice, originalSlice[i])
		}
	}
	return newSlice
}

func addElements(originalSlice []int, element int) []int {
	var newSlice []int
	newSlice = append(originalSlice, element)
	return newSlice
}

func copySlice(originalSlice []int) []int {
	var copiedSlice []int
	copiedSlice = make([]int, len(originalSlice))
	copy(copiedSlice, originalSlice)
	return copiedSlice
}

func removeElements(originalSlice []int, index int) []int {
	var newSlice []int
	if index < 0 || index >= len(originalSlice) {
		fmt.Println("Недоступный индекс")
		return originalSlice
	}
	for i := 0; i < len(originalSlice); i++ {
		if i != index {
			newSlice = append(newSlice, originalSlice[i])
		}
	}
	return newSlice
}

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}
	copied := copySlice(originalSlice)
	fmt.Println("Оригинальный слайс:", originalSlice)
	fmt.Println("Четные числа:", sliceExample(originalSlice))
	fmt.Println("Слайс после добавления элемента:", addElements(originalSlice, 3))
	fmt.Println("Копия слайса", copied)
	fmt.Println("Слайс после удаления элемента:", removeElements(originalSlice, 3))
	originalSlice[0] = 123
	fmt.Println("Изменённый оригинал:", originalSlice)
	fmt.Println("Копия после изменения оригинала:", copied)
}
