package main

import (
	"fmt"
	feature1 "study2/Feature1"
	"time"
)

func main() {

	fmt.Println("Саламалейкум")
	feature1.Feature1()
	for i := 0; i < 5; i++ {
		fmt.Println("Как же многа  строк")
		time.Sleep(1 * time.Second)
	}
}
