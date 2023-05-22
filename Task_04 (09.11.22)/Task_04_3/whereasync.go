/*Задача 3.
Реализуйте функцию where() (см. ДЗ №2, задача 2) в виде обобщённой функции с распараллеливанием.
Запустите 8 горутин, каждая из которых фильтрует свою часть исходного среза. Результаты фильтрации
затем объединяются в итоговый срез.
Важно: результаты работы обычной функции where() и функции where() с горутинами могут отличаться!
Это будут срезы из одних и тех же элементов, но взаимный порядок этих элементов может быть другим.
Проведите замеры производительности – есть ли выгода от использования горутин? При каком размере
среза-аргумента такая выгода начинает проявляться?
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func WhereParallel[T any](slice []T, predicate func(T) bool, threads int) (result []T) {
	length := len(slice)
	if threads < 1 || length <= threads {
		threads = 1
	}
	filtered := make(chan []T, threads)
	for i := 0; i < threads; i++ {
		from, to := i*(length/threads), (i+1)*(length/threads)
		switch {
		case i != threads-1:
			go whereSliceToChan(slice[from:to], predicate, filtered)
		case i == threads-1:
			go whereSliceToChan(slice[from:], predicate, filtered)
		}
	}
	for i := 0; i < threads; i++ {
		result = append(result, <-filtered...)
	}
	return
}

func whereSliceToChan[T any](slice []T, predicate func(T) bool, output chan []T) {
	var temp []T
	for _, val := range slice {
		if predicate(val) {
			temp = append(temp, val)
		}
	}
	output <- temp
}

func timeDecorator[T any](f func([]T, func(T) bool, int) []T) func([]T, func(T) bool, int) []T {
	return func(x []T, pr func(T) bool, thr int) []T {
		start := time.Now()
		y := f(x, pr, thr)
		fmt.Println(time.Since(start))
		return y
	}
}

func main() {

	const length int = 40000000
	const threads int = 100
	type dataType = int64
	var generator func(int64) dataType = rand.Int63n

	var predicate func(dataType) bool = func(x dataType) bool {
		return x%3 == 0 && x%5 == 2
	}

	rand.Seed(time.Now().Unix())
	data := make([]dataType, length)
	for i := 0; i < length; i++ {
		data[i] = generator(100)
	}

	slice1 := timeDecorator(WhereParallel[dataType])(data, predicate, threads)
	slice2 := timeDecorator(WhereParallel[dataType])(data, predicate, 1)
	fmt.Println(" slice1 Length : ", len(slice1))
	fmt.Println(" slice2 Length : ", len(slice2))
}
