package main

import "fmt"

// целых чисел. Нулевое значение представляет пустое множество.
type IntSet struct {
	words []uint64
}

func NewInt() *IntSet {
	newInt := IntSet{words: make([]uint64, 0)}
	return &newInt
}

// Has указывает, содержит ли множество неотрицательное значение х.
func (s *IntSet) Has(x int) bool {
	// fmt.Println(1 << 1)
	// fmt.Println(2 << 2)
	// fmt.Println(2 << 1)
	// fmt.Println(1 << 2)

	// fmt.Println(3 << 1)
	// fmt.Println(1 << 3)

	// fmt.Println(5 << 3)
	// 5 * 2^3 = 5 * 8 = 40
	// 5 * 2 = 10 10 * 2 = 20 20 * 2 = 40

	fmt.Println("1&1: ", 1&1)
	fmt.Println("2&1: ", 6&1)
	fmt.Println("1&2: ", 3&6)
	fmt.Println("2&2: ", 6&6)

	word, bit := x/64, uint(x%64)

	fmt.Println("number: ", x)
	fmt.Println("word: ", word)
	fmt.Println("bit: ", bit)
	fmt.Println("s.words[word]: ", s.words[word])
	fmt.Println("(1<<bit): ", (1 << bit))
	fmt.Println("s.words[word]&(1<<bit): ", s.words[word]&(1<<bit))

	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество,
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	x, y := NewInt(), NewInt()

	x.Add(34)
	x.Add(1872)
	x.Add(2)

	x.Has(34)
	x.Has(5)

	fmt.Println(x)

	y.Add(1)
	y.Add(89)
	y.Add(123)
	fmt.Println(y)

	x.UnionWith(y)

	fmt.Println(y)
	fmt.Println(x)
}
