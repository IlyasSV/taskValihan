package model

type Person struct {
	Name    string
	Age     int
	Address Address
}

//Измените структуру Person, чтобы она включала Address вместо строки Address.

//Задача 12: Указатели на структуры
//Создайте функцию IncrementAge, которая принимает указатель на Person и увеличивает возраст на 1.

func IncrementAge(a *Person) {
	a.Age += 1
}
