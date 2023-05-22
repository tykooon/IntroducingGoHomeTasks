/*
Задача 1.
Ваша задача – создать обобщённый тип для Стека. Используйте реализацию стека при помощи односвязного
списка (см. Training Center - Data Structures and Algorithms, презентация №2). Снабдите тип
для стека методами Push(), Pop() и IsEmpty(). Сделайте так, чтобы метод Pop() возвращал пару –
«значение-ошибка». Тип ошибки – стандартный интерфейс error (используйте пакет errors).
Ошибка возвращается, когда мы пытаемся извлечь элемент из пустого стека.
*/

package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	item     T // внутренние поля, недоступные вне пакета
	previous *Stack[T]
}

func NewStack[T any]() *Stack[T] {
	res := Stack[T]{previous: nil}
	return &res
}

func (stack Stack[T]) IsEmpty() bool {
	return stack.previous == nil
}

func (stack *Stack[T]) Push(item T) {
	prevItem := Stack[T]{item: item, previous: stack.previous}
	stack.previous = &prevItem
}

func (stack *Stack[T]) Pop() (T, error) {
	var res T
	if stack.IsEmpty() {
		return res, errors.New("stack is empty")
	}
	res = stack.previous.item
	stack.previous = stack.previous.previous
	return res, nil
}

func main() {
	intStack := *NewStack[int]()
	fmt.Println(intStack.IsEmpty())
	intStack.Push(28)
	intStack.Push(35)
	intStack.Push(923)
	intStack.Push(1223)
	fmt.Println(intStack.IsEmpty())
	for !intStack.IsEmpty() {
		item, err := intStack.Pop()
		if err == nil {
			fmt.Println(item, "  Ok!")
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(intStack.IsEmpty())
	_, err := intStack.Pop()
	if err != nil {
		fmt.Println(err.Error())
	}
}
