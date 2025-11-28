package main

type Semaphore chan struct{}

func (s Semaphore) Increment(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s Semaphore) Decrement(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	const count = 5

	var s = make(Semaphore, count)

	for i := 0; i < count; i++ {
		go func(n int) {
			defer s.Increment(1)

			print(n, " ")
		}(i)
	}

	s.Decrement(count)
}
