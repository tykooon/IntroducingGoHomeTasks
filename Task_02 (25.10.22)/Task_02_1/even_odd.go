/*Задача 1.
Функция получает на вход целое число (тип int) и возвращает пару значений.
Первое значение в паре – строка "even" или "odd", в зависимости от
количества простых делителей входного целого числа. То есть
15 = 3 * 5 -> e+ven, 30 = 2 * 3 * 5 -> odd, 9 = 3 * 3 -> even, 7 = 7 -> odd.
Второе значение в паре – булевский флаг. Он равен true, если входное целое
число больше или равно 2, и false, если входное число меньше или равно единице
(в этом случае первое значение в паре – пустая строка). В основной функции
main протестировать созданную функцию на нескольких входных данных.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkDivisors(-34))
	fmt.Println(checkDivisors(4))
	fmt.Println(checkDivisors(2))
	fmt.Println(checkDivisors(144))
	fmt.Println(checkDivisors(344521))
	fmt.Println(checkDivisors(4293423))
}

func checkDivisors(x int) (answer string, flag bool) {
	if x <= 1 {
		return "", false
	}
	x, counter := reduce(x, 2)
	x, temp := reduce(x, 3)
	counter += temp

	for i := 6; x > 1; i += 6 {
		var temp int
		x, temp = reduce(x, i-1)
		counter += temp
		if x == 1 {
			break
		}
		x, temp = reduce(x, i+1)
		counter += temp
	}
	if counter%2 == 0 {
		return "even", true
	} else {
		return "odd", true
	}
}

func reduce(num, div int) (int, counter int) {
	for ; num%div == 0; num /= div {
		counter++
	}
	return num, counter
}
