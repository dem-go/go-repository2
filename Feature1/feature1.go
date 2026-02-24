package feature1

import "fmt"

type people struct {
	name   string
	adress string
	age    int
}

func Feature1() {

	fmt.Println("Я первая фича")

	People := people{
		name:   "ДаНИл колбасенко",
		adress: "ДОм чикатило",
		age:    11,
	}
	fmt.Println("Инфа о человеке: ")
	fmt.Println("Имя: ", People.name)
	fmt.Println("Адресс: ", People.adress)
	fmt.Println("Возраст: ", People.age)
}
