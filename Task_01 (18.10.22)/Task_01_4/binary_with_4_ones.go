/*
	Задача 4.

Программа запрашивает у пользователя два целых числа a и b.
Затем программа выводит все положительные целые числа в диапазоне
от a (включительно) до b (включительно), которые в своём двоичном
представлении имеют ровно 4 единицы.
*/
package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	fmt.Println("Enter two limits of span (number1 [space] number2)")
	fmt.Scanln(&number1, &number2)

	if number1 > number2 {
		number1, number2 = number2, number1
	}

	for i := number1; i <= number2; i++ {
		var ones int
		for j := i; j > 0; j /= 2 {
			ones += j % 2
		}
		if ones == 4 {
			fmt.Println(i)
		}
	}
}
