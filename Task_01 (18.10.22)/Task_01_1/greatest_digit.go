/* Задача 1
Программа eзапрашивает у пользователя натуральное число, меньшее 1_000_000_000.
Если пользователь ввёл число не из указанного диапазона, программа выводит строку
«Input error» и останавливается. Иначе выводится самая большая цифра введённого числа.*/

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter natural number less than 1_000_000_000")
	fmt.Scanln(&number)

	if number <= 0 || number >= 1_000_000_000 {
		fmt.Println("Input error")
	} else {
		var digit int
		for number != 0 {
			temp := number % 10
			if temp > digit {
				digit = temp
			}
			number = (number - temp) / 10
		}
		fmt.Println("The greatest digit is: ", digit)
	}
}
