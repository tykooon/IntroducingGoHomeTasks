/*Задача 4.
Опишите интерфейс Cloner с методом клонирования Clone(). Создайте две-три простые структуры,
которые неявно реализуют интерфейс Cloner.
Создайте функцию sliceClone(). Её входной параметр – срез произвольных значений. Функция создаёт
и возвращает новый срез, состоящий из клонов значений из входного среза. Считаем, что
клонировать можно все базовые числовые типы, булевский тип, строку, а также структуры,
реализующие интерфейс Cloner(). То, что нельзя клонировать, в срез-результат не попадает.
*/

package main

import "fmt"

type Cloner interface {
	Clone() any
}

type Client struct {
	Id   int
	Name string
}

func (c *Client) Clone() (res any) {
	if c != nil {
		res = Client{c.Id, c.Name}
	}
	return
}

type Product struct {
	Id    int
	Name  string
	Price int
}

func (p *Product) Clone() (res any) {
	if p != nil {
		res = Product{p.Id, p.Name, p.Price}
	}
	return
}

type Order struct {
	Id       int
	Client   Client
	Products map[Product]int
}

func (o *Order) Clone() any {
	if o == nil {
		return nil
	}
	res := Order{o.Id, o.Client, make(map[Product]int)}
	for k, v := range o.Products {
		res.Products[k] = v
	}
	return res
}

func (o *Order) TotalPrice() (result int) {
	for k, v := range o.Products {
		result += v * k.Price
	}
	return
}

func (o *Order) AddToCart(product Product, amount int) {
	if amount <= 0 {
		return
	}
	_, isInCart := o.Products[product]
	if !isInCart {
		o.Products[product] = 0
	}
	o.Products[product] += amount
}

func sliceClone(slice []any) (result []any) {
	if len(slice) == 0 {
		return
	}
	for i := range slice {
		switch slice[i].(type) {
		case Cloner:
			result = append(result, slice[i].(Cloner).Clone())
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
			float32, float64, string, bool:
			result = append(result, slice[i])
		}
	}
	return
}

func main() {
	alex := Client{1, "Alex"}
	john, ok := alex.Clone().(Client)
	if ok {
		john.Id++
		john.Name = "John Doe"
	}
	kindle := Product{1, "Kindle Paperwhite 2", 130}
	xBox := Product{2, "Microsoft X-Box Series S", 500}
	gamepad := Product{3, "Canyon Gamepad for MS XBox X/S", 40}

	order1 := Order{1, alex, make(map[Product]int)}
	order1.AddToCart(kindle, 2)
	order1.AddToCart(xBox, 1)

	order2 := Order{2, john, make(map[Product]int)}
	order2.AddToCart(xBox, 1)
	order2.AddToCart(gamepad, 2)
	order2.AddToCart(kindle, -2)
	order3, ok := order1.Clone().(Order)
	if !ok {
		fmt.Println("Something's gone wrong!")
		return
	}
	order3.Id = 3
	order1.AddToCart(gamepad, 1)

	fmt.Println("Order 1 Total Price: ", order1.TotalPrice())
	fmt.Println("Order 3 Total Price: ", order3.TotalPrice())
	fmt.Println("Order 2 Total Price: ", order2.TotalPrice())

	var anySlice []any
	anySlice = append(anySlice, &kindle, &alex, order1.AddToCart, true, "John Doe")
	anySlice = append(anySlice, struct {
		Name string
		Age  int
	}{Name: "Nick", Age: 25})

	temp := sliceClone(anySlice)
	for i := range temp {
		fmt.Printf("%T\n", temp[i])
	}
}
