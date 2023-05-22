/*Задача 3.
Создать функцию sequence(), получающую на вход произвольное количество целых чисел.
Логика работы функции зависит от количества её аргументов:
•	одно число a: вернуть срез с целыми числами от 0 до a включительно (или от a до 0, если a – отрицательное);
•	два числа a и b: вернуть срез с целыми числами от a до b (или от b до a, если b < a);
•	чисел больше двух – вернуть срез с этими числами;
•	нет чисел – вернуть пустой срез.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(sequence(1, 2, 3, 4, 6))
	fmt.Println(sequence(5, -4))
	fmt.Println(sequence(-7))
	fmt.Println(sequence(0))
	fmt.Println(sequence())
	fmt.Println(append(sequence(), 23))
}

func sequence(num ...int) (result []int) {
	switch len(num) {
	case 0:
		return
	case 1:
		num = append(num, 0)
		fallthrough
	case 2:
		if num[1] < num[0] {
			num[0], num[1] = num[1], num[0]
		}
		for ; num[0] <= num[1]; num[0]++ {
			result = append(result, num[0])
		}
		return
	default:
		return num
	}
}
