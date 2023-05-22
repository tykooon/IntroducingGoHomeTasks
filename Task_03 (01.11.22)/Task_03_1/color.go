/*Задача 1.
Объявите тип Color, использовав в качестве базового типа массив из трёх байт.
Считаем, что экземпляр типа Color представляет цвет в формате RGB. Опишите
для типа Color следующие методы: Print() выводит информацию цвета на консоль,
GetR(), GetG(), GetB() возвращают значения компонент цвета (тип возвращаемого
значения – byte). SetR(), SetG(), SetB() позволяют установить компоненты цвета.
Метод GetBrightness() возвращает яркость цвета, вычисленную по формуле
(0.2126*R + 0.7152*G + 0.0722*B). Создайте функцию maxBrightness() – она
получает на вход срез значений Color, а возвращает указатель на элемент из среза,
имеющий максимальную яркость.
*/

package main

import (
	"fmt"
)

type Color [3]byte

func (c Color) Print() {
	fmt.Printf("R: %d G: %d B: %d", c[0], c[1], c[2])
	fmt.Println()
}

func (c *Color) GetR() byte {
	return c[0]
}

func (c *Color) GetG() byte {
	return c[1]
}

func (c *Color) GetB() byte {
	return c[2]
}

func (c *Color) SetR(r byte) {
	c[0] = r
}

func (c *Color) SetG(g byte) {
	c[1] = g
}

func (c *Color) SetB(b byte) {
	c[2] = b
}

func (c *Color) GetBrightness() float64 {
	return float64(c[0])*0.2126 + float64(c[1])*0.7152 + float64(c[2])*0.0722
}

func maxBrightness(colorSlice []Color) (result *Color) {
	if len(colorSlice) == 0 {
		return nil
	}
	result = &colorSlice[0]
	var brightness = result.GetBrightness()
	for i, color := range colorSlice {
		var temp = color.GetBrightness()
		if temp > brightness {
			result = &colorSlice[i]
			brightness = temp
		}
	}
	return
}

func main() {
	color := Color([3]byte{0x10, 0x40, 0xff})
	color.Print()

	slice := []Color{
		color,
		Color([3]byte{0x10, 0x40, 0xff}),
		Color([3]byte{0xf0, 0xf0, 0xf0}),
		Color([3]byte{0x00, 0xff, 0xff}),
		Color([3]byte{0xf0, 0x40, 0xd0}),
	}

	slice[2].SetR(254) // Без этой строчки самый яркий цвет в массиве - с индексом 1.

	maxBrightness(slice).Print()
}
