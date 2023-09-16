// Если структуры включены в общую и имеют одинаковые методы, то будет ошибка компиляции.
// Встраиваемые структуры могут перекрывать методы, если встраиваемая структура
// и включающая структура имеют методы с одинаковыми именами.

package main

import "fmt"

type Animals struct {
	Dog
	Cat
}

func (a *Animals) Speak() {
	fmt.Println("I am an animal.")
}

type Dog struct {
}

func (a *Dog) Speak() {
	fmt.Println("Dog.")
}

func (d *Cat) Speak() {
	fmt.Println("Cat!")
}

type Cat struct {
}

func main() {
	animal := Animals{}
	animal.Speak() // Выводит "Woof!"
}
