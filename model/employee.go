package model

import "fmt"

//Задача 9: Анонимные поля
//Создайте структуру Employee, которая будет включать анонимное поле типа Person,
// а также поле Position.

type Employee struct {
	Person
	Position int
}

//Задача 11: Методы для структур с анонимными полями
//Добавьте метод Describe для структуры Employee, который будет выводить информацию о сотруднике,
//включая данные из анонимного поля Person.

func (e Employee) Describe() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %+v\nPosition: %d\n",
		e.Name, e.Age, e.Address, e.Position)
}
