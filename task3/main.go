package main

import "taskValihan1/model"

func main() {
	emp := model.Employee{
		Person: model.Person{
			Name: "111",
			Age:  20,
			Address: model.Address{
				City:       "498",
				Street:     "2342",
				PostalCode: 23235,
			},
		},
		Position: 1,
	}

	emp.Describe()

	emps := []model.Employee{
		{Person: model.Person{
			Name: "123",
			Age:  29,
			Address: model.Address{
				City:       "1244",
				Street:     "234525",
				PostalCode: 12414,
			},
		},
			Position: 2,
		},
		{Person: model.Person{
			Name: "235",
			Age:  34,
			Address: model.Address{
				City:       "235",
				Street:     "324462",
				PostalCode: 235034,
			},
		},
			Position: 3,
		},
		{Person: model.Person{
			Name: "2534",
			Age:  23,
			Address: model.Address{
				City:       "dfsdf",
				Street:     "afsf",
				PostalCode: 2335,
			},
		},
			Position: 4,
		},
	}

	for _, element := range emps {
		element.Describe()
	}

}
