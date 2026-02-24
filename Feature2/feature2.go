package feature2

import (
	"fmt"
	"time"
)

func mine(coal int) {
	for i := 0; i < 5; i++ {
		s := i
		coal = coal + s
		fmt.Println("уголь: ", coal)
	}
}

func Feature2() {

	fmt.Println("Вот они ваши эти хвалённые горутины: ")
	go func() {
		fmt.Println("Я горутина")
	}()
	go mine(0)
	time.Sleep(500 * time.Millisecond)
}
