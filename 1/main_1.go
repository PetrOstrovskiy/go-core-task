package main

import (
	"crypto/sha256"
	"fmt"
)

// 3
func buildString(numDecimal, numOctal, numHexadecimal int,
	pi float64, name string, isActive bool, complexNum complex64) string {

	return fmt.Sprintf("%d-%o-%x-%.2f-%s-%t-%v",
		numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
}

// 4
func toRuneSlice(s string) []rune {
	return []rune(s)
}

// 5
func hashWithSalt(runes []rune, salt string) [32]byte {
	mid := len(runes) / 2
	saltRunes := []rune(salt)

	result := append([]rune{}, runes[:mid]...)
	result = append(result, saltRunes...)
	result = append(result, runes[mid:]...)

	return sha256.Sum256([]byte(string(result)))
}

func main() {
	// 1
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	// 2
	fmt.Printf("Тип переменной numDecimal: %T\n", numDecimal)
	fmt.Printf("Тип переменной numOctal: %T\n", numOctal)
	fmt.Printf("Тип переменной numHexadecimal: %T\n", numHexadecimal)
	fmt.Printf("Тип переменной pi: %T\n", pi)
	fmt.Printf("Тип переменной name: %T\n", name)
	fmt.Printf("Тип переменной isActive: %T\n", isActive)
	fmt.Printf("Тип переменной complexNum: %T\n", complexNum)

	// 3
	result := buildString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\nСтрока:", result)

	// 4
	runes := toRuneSlice(result)
	fmt.Println("\nСрез рун:", runes)
	fmt.Println("Строка:", string(runes))

	// 5
	hash := hashWithSalt(runes, "go-2024")
	fmt.Println("\nSHA256 с солью:")
	fmt.Printf("%x\n", hash)
}
