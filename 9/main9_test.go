package main

import (
	"reflect"
	"testing"
)

func TestToFloat(t *testing.T) {
	in := make(chan uint8)

	go func() {
		in <- 1
		in <- 5
		in <- 10
		close(in)
	}()

	out := toFloat(in)

	var result []float64
	for v := range out {
		result = append(result, v)
	}

	expected := []float64{1, 5, 10}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCube(t *testing.T) {
	in := make(chan float64)

	go func() {
		in <- 2
		in <- 3
		in <- 4
		close(in)
	}()

	out := cube(in)

	var result []float64
	for v := range out {
		result = append(result, v)
	}

	expected := []float64{8, 27, 64}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFullPipeline(t *testing.T) {

	in := make(chan uint8)
	go func() {
		for _, v := range []uint8{2, 3, 5} {
			in <- v
		}
		close(in)
	}()

	stage1 := toFloat(in)
	stage2 := cube(stage1)

	var result []float64
	for v := range stage2 {
		result = append(result, v)
	}

	expected := []float64{8, 27, 125}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestEmptyInput(t *testing.T) {
	in := make(chan uint8)
	close(in)

	out1 := toFloat(in)
	out2 := cube(out1)

	var result []float64
	for v := range out2 {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}
