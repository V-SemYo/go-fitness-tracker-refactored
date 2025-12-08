package personaldata

import "fmt"

type Personal struct {
	// Personal содержит персональные данные пользователя
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// Print выводит данные пользователя в консоль
	fmt.Println("Имя: ", p.Name)
	fmt.Println("Вес: ", p.Weight)
	fmt.Println("Рост: ", p.Height)
}
