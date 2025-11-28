package main

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func makeTestChannel(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func TestJoinChannelsBasic(t *testing.T) {
	ch1 := makeTestChannel([]int{1, 2, 3})
	ch2 := makeTestChannel([]int{10, 20})
	ch3 := makeTestChannel([]int{100})

	merged := joinChannels(ch1, ch2, ch3)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 10, 20, 100}

	sort.Ints(result)
	sort.Ints(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestJoinChannelsEmpty(t *testing.T) {
	// Все каналы пустые
	ch1 := makeTestChannel([]int{})
	ch2 := makeTestChannel([]int{})

	merged := joinChannels(ch1, ch2)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("Ожидался пустой срез, получено %v", result)
	}
}

func TestJoinChannelsAsync(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 100
		time.Sleep(5 * time.Millisecond)
		ch2 <- 200
		close(ch2)
	}()

	merged := joinChannels(ch1, ch2)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	expected := []int{1, 2, 100, 200}
	sort.Ints(result)
	sort.Ints(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
