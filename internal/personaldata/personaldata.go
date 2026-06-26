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
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f кг.\n", p.Weight)
	fmt.Printf("Рост: %.2f м.\n", p.Height)
}
