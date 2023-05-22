/*Задача 2.
	Создать функцию where(), которая получает на вход срез целых чисел и функцию-предикат.
Функция-предикат получает на вход целое число и возвращает true/false. Функция where()
возвращает срез, состоящий из чисел среза-аргумента, удовлетворяющих функции-предикату.
	Создать функцию foreach() для среза целых чисел и функции-действия. Функция-действие
выполняет некий код для своего целого аргумента. Функция foreach() запускает
функцию-действие для каждого числа из своего среза-аргумента.
	Протестировать работу функций, используя в качестве предиката
и действия анонимные функции.
*/

package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 4, 6, 8, 9, 5, 4, 9, 6}
	fmt.Println(where(slice, func(x int) bool { return x%2 == 0 }))

	foreach(slice, func(x int) { fmt.Print(x*x, " ") })

	fmt.Println()
	foreach(where(slice, func(int) bool { return true }), func(int) { fmt.Print(string(0x1F60E)) })
}

func where(slice []int, predicate func(int) bool) (result []int) {
	for _, val := range slice {
		if predicate(val) {
			result = append(result, val)
		}
	}
	return
}

func foreach(slice []int, action func(int)) {
	for _, val := range slice {
		action(val)
	}
}
