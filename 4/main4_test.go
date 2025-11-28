package main

import (
	"reflect"
	"testing"
)

func TestResultSliceBasic(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := resultSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestResultSliceEmptySecond(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{}

	expected := []string{"a", "b", "c"}

	result := resultSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestResultSliceEmptyFirst(t *testing.T) {
	// Если первый слайс пустой — результат пустой
	slice1 := []string{}
	slice2 := []string{"a", "b"}

	expected := []string{}

	result := resultSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось пустой слайс, получено %v", result)
	}
}

func TestResultSliceNoDifferences(t *testing.T) {
	slice1 := []string{"x", "y"}
	slice2 := []string{"x", "y", "z"}

	expected := []string{}

	result := resultSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestResultSliceDuplicates(t *testing.T) {
	slice1 := []string{"a", "b", "a", "c"}
	slice2 := []string{"b"}

	expected := []string{"a", "a", "c"}

	result := resultSlice(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
