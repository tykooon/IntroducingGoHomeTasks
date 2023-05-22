/*Задача 2.
Проверить, является ли введённое положительное целое число
палиндромом, если записать его в троичной системе счисления.
*/

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter integer number")
	fmt.Scanln(&number)

	var inputNum = number
	var reverseNum int

	for number > 0 {
		var digit = number % 3
		reverseNum = reverseNum*3 + digit
		number = number / 3
	}

	if inputNum == reverseNum {
		fmt.Println(inputNum, " -- is palindromic in ternary system.")
	} else {
		fmt.Println(inputNum, " -- is not palindromic in ternary system.")
	}
}
