package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	// Проверка добавления пары "ключ-значение"
	m := NewStringIntMap()
	m.Add("a", 10)

	if val, ok := m.data["a"]; !ok || val != 10 {
		t.Errorf("Ожидалось, что ключ 'a' будет иметь значение 10, получено: %v, существует: %v", val, ok)
	}
}

func TestRemove(t *testing.T) {
	// Проверка удаления элемента по ключу
	m := NewStringIntMap()
	m.Add("a", 10)
	m.Remove("a")

	if _, ok := m.data["a"]; ok {
		t.Errorf("Ключ 'a' должен быть удалён из карты")
	}
}

func TestCopy(t *testing.T) {
	// Проверка корректного копирование карты
	m := NewStringIntMap()
	m.Add("a", 10)
	m.Add("b", 20)

	copied := m.Copy()

	if !reflect.DeepEqual(m.data, copied) {
		t.Errorf("Ожидалась точная копия карты. Получено %v, ожидалось %v", copied, m.data)
	}

	m.Add("c", 30)
	if _, ok := copied["c"]; ok {
		t.Errorf("Копия карты не должна изменяться после изменения оригинала")
	}
}

func TestExists(t *testing.T) {
	// Проверка наличия ключа
	m := NewStringIntMap()
	m.Add("a", 10)

	if !m.Exists("a") {
		t.Errorf("Ожидалось, что ключ 'a' существует, но Exists вернул false")
	}

	if m.Exists("b") {
		t.Errorf("Ожидалось, что ключ 'b' отсутствует, но Exists вернул true")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 10)

	val, ok := m.Get("a")
	if !ok || val != 10 {
		t.Errorf("Ожидалось получить значение 10 для ключа 'a', получено %v (exists=%v)", val, ok)
	}

	// Нет ключа
	_, ok = m.Get("b")
	if ok {
		t.Errorf("Get должен возвращать exists=false для отсутствующего ключа 'b'")
	}
}
