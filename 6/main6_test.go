package main

import "testing"

func TestRandNumsGenerator(t *testing.T) {
	n := 10
	ch := randNumsGenerator(n)

	count := 0
	for v := range ch {
		if v < 0 || v >= n {
			t.Errorf("Сгенерировано значение %d вне диапазона [0, %d)", v, n)
		}
		count++
	}

	if count != n {
		t.Errorf("Ожидалось %d чисел, получено %d", n, count)
	}
}

func TestRandNumsGeneratorEmpty(t *testing.T) {
	ch := randNumsGenerator(0)

	_, ok := <-ch
	if ok {
		t.Errorf("Канал должен быть сразу закрыт при n=0")
	}
}
