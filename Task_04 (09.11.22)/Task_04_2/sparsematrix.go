/*  Задача 2.
Разреженные числовые матрицы – числовые матрицы, в которых только относительно малая часть элементов
не равна нулю. Необходимо создать обобщённый тип для представления разреженной числовой матрицы.
Элементами матрицы могут быть стандартные числовые типы (кроме комплексных чисел) и типы,
созданные на их основе. Предусмотрите методы для чтения и записи элемента матрицы (элемент
определяется двумя индексами). В этой задаче при некорректности значений (например, индексы
выходят за допустимый диапазон) вызывайте панику.
*/

package main

import "fmt"

type SimpleNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float64
}

type SparseMatrix[T SimpleNumber] struct {
	rows     int
	cols     int
	elements map[int]T
}

func NewSparseMatrix[T SimpleNumber](rows, cols int) *SparseMatrix[T] {
	if rows <= 0 || cols <= 0 {
		panic("Wrong matrix sizes")
	}
	return &SparseMatrix[T]{rows: rows, cols: cols, elements: make(map[int]T)}
}

func (matr SparseMatrix[T]) GetSizes() (i, j int) {
	return matr.rows, matr.cols
}

func (matr SparseMatrix[T]) GetElement(i, j int) T {
	var res T
	if i < 0 || j < 0 || i >= matr.rows || j >= matr.cols {
		panic("Wrong indices")
	}
	k := i*matr.cols + j
	item, ok := matr.elements[k]
	if ok {
		res = item
	}
	return res
}

func (matr SparseMatrix[T]) SetElement(i, j int, value T) {
	var def T
	if i < 0 || j < 0 || i >= matr.rows || j >= matr.cols {
		panic("Wrong indices")
	}
	k := i*matr.cols + j
	_, ok := matr.elements[k]
	switch {
	case ok && value == def:
		delete(matr.elements, k)
	case !ok && value == def:
		break
	case (!ok && value != def) || (ok && value != def):
		matr.elements[k] = value
	}
}

func (matr SparseMatrix[T]) PrintEssential() {
	for k, v := range matr.elements {
		fmt.Println(k/matr.cols, ", ", k%matr.cols, " : ", v)
	}
}

func main() {
	spM := NewSparseMatrix[byte](20, 30)

	fmt.Println(spM.GetElement(4, 6))
	spM.SetElement(0, 0, 0xdb)
	spM.SetElement(4, 6, 0xad)
	spM.PrintEssential()

	fmt.Println("----------------------")
	spM.SetElement(0, 0, 0x98)
	spM.SetElement(2, 4, 0)
	spM.SetElement(4, 6, 0)
	spM.PrintEssential()
}
