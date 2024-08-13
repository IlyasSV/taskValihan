package model

//Задача 9: Анонимные поля
//Создайте структуру Employee, которая будет включать анонимное поле типа Person,
// а также поле Position.

type Employee struct {
	Person
	Position int
}
