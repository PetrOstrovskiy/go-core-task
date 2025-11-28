package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	result := sliceExample(original)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	original := []int{1, 2, 3}
	element := 10
	expected := []int{1, 2, 3, 10}

	result := addElements(original, element)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{5, 10, 15}
	copied := copySlice(original)

	if !reflect.DeepEqual(original, copied) {
		t.Errorf("Слайсы должны быть равны. Оригинал %v, копия %v", original, copied)
	}

	original[0] = 99
	if copied[0] == 99 {
		t.Errorf("Копия изменилась после изменения оригинала!")
	}
}

func TestRemoveElements(t *testing.T) {
	original := []int{10, 20, 30, 40, 50}
	index := 2
	expected := []int{10, 20, 40, 50}

	result := removeElements(original, index)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestRemoveElementsInvalidIndex(t *testing.T) {
	original := []int{1, 2, 3}

	// отрицательный индекс
	result := removeElements(original, -1)
	if !reflect.DeepEqual(result, original) {
		t.Errorf("При неверном индексе (−1) должен возвращаться оригинальный слайс")
	}

	// индекс за границей
	result = removeElements(original, 5)
	if !reflect.DeepEqual(result, original) {
		t.Errorf("При неверном индексе (5) должен возвращаться оригинальный слайс")
	}
}
