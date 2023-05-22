/*Задача 2.
Создайте структуру Matrix для представления двумерной матрицы чисел (элементы матрицы имеют	тип
float64). Поля структуры: rows – число строк матрицы (больше нуля), cols – число столбцов матрицы
(больше нуля), data – одномерный срез с элементами матрицы.
Опишите функцию-конструктор для создания экземпляра матрицы. Параметры функции – число строк и
число столбцов. В случае некорректных значений параметров создавайте матрицу размером 1х1.
Опишите два метода для структуры Matrix. Метод Get(i, j) возвращает элемент матрицы с индексами
i,j (индексация начинается с нуля). Метод Set(i, j, value) устанавливает элемент матрицы с индексами
i,j. Обратите внимание, что методы должны транслировать два входных индекса в один индекс в срезе
с элементами матрицы. Если индексы выходят за допустимый диапазон, метод Get(i, j) возвращает 0,
а метод Set(i, j, value) ничего не делает.
Создайте метод Print() для структуры Matrix, который выводит матрицу на консоль.
*/

package main

import (
	"fmt"
)

type Matrix struct {
	rows int
	cols int
	data []float64
}

func NewMatrix(rows, cols int) (result Matrix) {
	if cols <= 0 || rows <= 0 {
		cols, rows = 1, 1
	}
	result.rows, result.cols = rows, cols
	result.data = make([]float64, cols*rows)
	return
}

func (matr *Matrix) Get(i, j int) float64 {
	if i < 0 || j < 0 || i >= matr.rows || j >= matr.cols {
		return 0
	}
	return matr.data[i*matr.cols+j]
}

func (matr *Matrix) Set(i, j int, value float64) {
	if i >= 0 && j >= 0 && i < matr.rows && j < matr.cols {
		matr.data[i*matr.cols+j] = value
	}
}

func (matr *Matrix) Print() {
	for i := 0; i < matr.rows; i++ {
		rowStart := i * matr.cols
		for j := 0; j < matr.cols; j++ {
			fmt.Printf(" %.3f ", matr.data[rowStart+j])
		}
		fmt.Println()
	}
}

func (matr *Matrix) SizeInfo() {
	fmt.Printf("Matrix size = %d x %d", matr.rows, matr.cols)
	fmt.Println()
}

func main() {
	fmt.Println("Ready!")

	matrix := NewMatrix(2, 3)
	matrix.SizeInfo()

	matrix.Set(0, 1, 2.2)
	matrix.Set(1, 2, 4.5)

	matrix.Print()

	badMatrix := NewMatrix(-2, 0)
	badMatrix.SizeInfo()
	fmt.Println("badMatrix(0,0) = ", badMatrix.Get(0, 0))

}
