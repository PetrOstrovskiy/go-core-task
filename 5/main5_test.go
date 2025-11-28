package main

import (
	"reflect"
	"testing"
)

func TestSameElementsBasic(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	expectedBool := true
	expectedSlice := []int{64, 3}

	has, result := sameElements(a, b)

	if has != expectedBool {
		t.Errorf("Ожидалось true, получено %v", has)
	}

	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Ожидалось %v, получено %v", expectedSlice, result)
	}
}

func TestSameElementsNoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	expectedBool := false

	has, result := sameElements(a, b)

	if has != expectedBool {
		t.Errorf("Ожидалось false, получено %v", has)
	}

	if result != nil && len(result) != 0 {
		t.Errorf("Ожидался пустой срез или nil, получено %v", result)
	}
}

func TestSameElementsWithDuplicates(t *testing.T) {
	a := []int{1, 2, 2, 3}
	b := []int{2, 2, 2}

	expectedBool := true
	expectedSlice := []int{2, 2}

	has, result := sameElements(a, b)

	if has != expectedBool {
		t.Errorf("Ожидалось true, получено %v", has)
	}

	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Ожидалось %v, получено %v", expectedSlice, result)
	}
}

func TestSameElementsEmptyA(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}

	has, result := sameElements(a, b)

	if has != false {
		t.Errorf("Ожидалось false, получено %v", has)
	}

	if result != nil && len(result) != 0 {
		t.Errorf("Ожидался пустой результат, получено %v", result)
	}
}

func TestSameElementsEmptyB(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{}

	has, result := sameElements(a, b)

	if has != false {
		t.Errorf("Ожидалось false, получено %v", has)
	}

	if result != nil && len(result) != 0 {
		t.Errorf("Ожидался пустой результат, получено %v", result)
	}
}
