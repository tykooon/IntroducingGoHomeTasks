/*Задача 4.
Функция group() получает в качестве аргумента карту, в которой ключ имеет тип byte,
а значение для ключа – тип string. Результат работы функции – карта с байтовыми ключами
из диапазона 0...9 и значениями для ключа в виде срезов из строк.
Функция group() выполняет группировку значений из исходной карты по признаку последней
цифры числа-ключа. Например, если в исходной карте есть пары 11:"red" и 51:"green",
то они должны быть сгруппированы в пару 1:["red", "green"]. Пар с пустыми срезами в
карте-результате быть не должно.
*/

package main

import (
	"fmt"
)

func main() {
	mapTest := map[byte]string{
		12:  "twelve",
		22:  "twenty two",
		14:  "fourteen",
		4:   "four",
		17:  "seventeen",
		37:  "thirty seven",
		127: "one hundred twenty seven",
	}
	fmt.Println(group(mapTest))
}

func group(mapIn map[byte]string) map[byte][]string {
	mapOut := make(map[byte][]string, 10)
	for key, val := range mapIn {
		tmpKey := key % 10
		if _, isExists := mapOut[tmpKey]; !isExists {
			mapOut[tmpKey] = make([]string, 0)
		}
		mapOut[tmpKey] = append(mapOut[tmpKey], val)
	}
	return mapOut
}
