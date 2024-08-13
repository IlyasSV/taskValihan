package main

import (
	"fmt"
	"taskValihan1/model"
)

//Задача 12: Указатели на структуры
//Создайте функцию IncrementAge, которая принимает указатель на Person и увеличивает возраст на 1.

func main() {
	person := &model.Person{
		Name: "REFs",
		Age:  24,
		Address: model.Address{
			City:       "rsfgs",
			Street:     "sfsgdg",
			PostalCode: 1243,
		},
	}

	model.IncrementAge(person)
	fmt.Printf("%+v\n", person)
}
