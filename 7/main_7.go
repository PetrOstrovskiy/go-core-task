package main

import (
	"fmt"
	"sync"
)

func joinChannels(channels ...<-chan int) <-chan int {
	resultChannel := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(len(channels))

	for _, channel := range channels {
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				resultChannel <- value
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	return resultChannel
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ch2 <- i + 50
		}
		close(ch2)
	}()

	for num := range joinChannels(ch1, ch2) {
		fmt.Println(num)
	}

}
