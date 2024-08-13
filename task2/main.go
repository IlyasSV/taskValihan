package main

import (
	"fmt"
	"taskValihan1/model"
)

// Задача 8: Использование структуры с композицией
// Инициализируйте несколько экземпляров структуры Person с использованием новой структуры Address
// и выведите их на экран.

func main() {
	person := []model.Person{
		{Name: "44", Age: 24, Address: model.Address{
			City:       "4141",
			Street:     "235",
			PostalCode: 4444,
		}},
		{Name: "55", Age: 25, Address: model.Address{
			City:       "5252",
			Street:     "5757",
			PostalCode: 5959,
		}},
		{Name: "66", Age: 27, Address: model.Address{
			City:       "5252",
			Street:     "75757",
			PostalCode: 5959,
		}},
	}

	fmt.Printf("%+v\n", person)
}
