package main

import "fmt"

func toFloat(in <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for v := range in {
			out <- float64(v)
		}
	}()
	return out
}

func cube(in <-chan float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v * v
		}
	}()
	return out
}

func main() {
	in := make(chan uint8)

	go func() {
		for i := uint8(0); i <= 10; i++ {
			in <- i
		}
		close(in)
	}()

	floatNum := toFloat(in)
	cubeNum := cube(floatNum)

	for v := range cubeNum {
		fmt.Println(v)
	}
}
