package main

import (
	"fmt"
	"taskValihan1/model"
)

//Задача 10: Использование анонимных полей
//Инициализируйте несколько экземпляров структуры Employee и выведите их на экран.

func main() {
	emp := []model.Employee{
		{Person: model.Person{Name: "11", Age: 20},
			Position: 1},
		{Person: model.Person{Name: "22", Age: 21},
			Position: 2},
		{Person: model.Person{Name: "33", Age: 22},
			Position: 3},
	}
	fmt.Printf("%+v\n", emp)

}
