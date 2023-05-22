/*Задача 3.
Опишите структуру Point для представления точки на плоскости (две координаты X и Y – вещественные
числа). Создайте метод для структуры Point для установки новых координат точки (изменения X и Y).
Создайте структуру PointLabeled – это точка на плоскости, снабжённая текстовой меткой.
Создайте функцию normalize(). Она должна получать на вход срез из точек, причём это могут быть
как экземпляры структуры Point, так и экземпляры структуры PointLabeled (причём даже вперемешку).
Функция нормализует координаты всех точек в срезе – делает так, чтобы координаты вписывались
в единичный квадрат [0, 1]x[0, 1]. Так, минимальная координата X всех точек становиться нулём,
максимальная – 1, а остальные изменяются пропорционально (https://en.wikipedia.org/wiki/Feature_scaling).
В этой задаче вы можете самостоятельно создавать необходимые (или полезные) для решения задачи интерфейсы и методы.
*/

package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p *Point) SetX(x float64) {
	p.X = x
}

func (p *Point) SetY(y float64) {
	p.Y = y
}

func (p *Point) SetXY(x, y float64) {
	p.SetX(x)
	p.SetY(y)
}

func NewPoint(x, y float64) (result Point) {
	result.SetXY(x, y)
	return
}

func (p *Point) GetXY() (float64, float64) {
	return p.X, p.Y
}

type PointLabeled struct {
	Point Point
	Label string
}

func (pl *PointLabeled) SetXY(x, y float64) {
	pl.Point.SetX(x)
	pl.Point.SetY(y)
}

func (pl *PointLabeled) SetLabel(label string) {
	pl.Label = label
}

func (pl *PointLabeled) GetXY() (float64, float64) {
	return pl.Point.X, pl.Point.Y
}

type PointHelper interface {
	SetXY(x, y float64)
	GetXY() (float64, float64)
}

func normalize(slice []PointHelper) []PointHelper {
	if len(slice) == 0 {
		return slice
	}
	xMin, yMin, xMax, yMax := findBorders(slice)
	dx := xMax - xMin
	dy := yMax - yMin
	if dx == 0 { //Возможно, правильнее указать dx < 5e-320  , например
		dx = 1
	}
	if dy == 0 {
		dy = 1
	}
	for _, point := range slice {
		x, y := point.GetXY()
		point.SetXY((x-xMin)/dx, (y-yMin)/dy)
	}
	return slice
}

func findBorders(slice []PointHelper) (xMin, yMin, xMax, yMax float64) {
	if len(slice) == 0 {
		return
	}
	xMin, yMin = (slice)[0].GetXY()
	xMax, yMax = (slice)[0].GetXY()
	for _, point := range slice {
		x, y := point.GetXY()
		xMin, xMax = shiftBorder(x, xMin, xMax)
		yMin, yMax = shiftBorder(y, yMin, yMax)
	}
	return
}

func shiftBorder(x, min, max float64) (float64, float64) {
	switch {
	case x > max:
		max = x
	case x < min:
		min = x
	}
	return min, max
}

func main() {

	p := &Point{10, .5}
	pL := &PointLabeled{Point{7, .2}, "p1"}
	p2 := &Point{4, .8}

	slice := []PointHelper{p, pL, p2}
	normalize(slice)
	for i := range slice {
		x, y := slice[i].GetXY()
		fmt.Printf("Point %d: (%.3f,%.3f)\n", i, x, y)
	}
}
