package main

import (
	"testing"
)

func TestBuildString(t *testing.T) {
	s := buildString(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
	expected := "42-52-2a-3.14-Golang-true-(1+2i)"

	if s != expected {
		t.Errorf("ожидалось %s, получено %s", expected, s)
	}
}

func TestRuneSlice(t *testing.T) {
	r := toRuneSlice("Go")
	if len(r) != 2 || r[0] != 'G' || r[1] != 'o' {
		t.Error("ошибка преобразования строки в руны")
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("test")
	hash1 := hashWithSalt(runes, "go-2024")
	hash2 := hashWithSalt(runes, "go-2024")

	if hash1 != hash2 {
		t.Error("хэш с одинаковыми данными должен совпадать")
	}
}
