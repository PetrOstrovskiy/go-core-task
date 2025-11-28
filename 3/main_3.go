package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range s.data {
		newMap[k] = v
	}
	return newMap
}

func (s *StringIntMap) Exists(key string) bool {
	_, exist := s.data[key]
	return exist
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, exist := s.data[key]
	return value, exist
}

func main() {
	m := NewStringIntMap()
	m.Add("a", 10)
	m.Add("b", 20)
	fmt.Println("Map:", m)
	m.Remove("a")
	fmt.Println("Map после удаления", m)
	copied := m.Copy()
	fmt.Println("Скопированная map:", copied)
	m.Add("c", 30)
	fmt.Println("Map", m)
	fmt.Println("Скопированная map:", copied)
	fmt.Println(m.Exists("a"))
	fmt.Println(m.Exists("b"))
	fmt.Println(m.Get("c"))
}
