package main

import (
	"context"
	"fmt"
	feature1 "study2/Feature1"
	feature2 "study2/Feature2"
	simple_connection "study2/feature_postgres/simple_connection"
	simple_sql "study2/feature_sql"
)

func main() {
	ctx := context.Background()

	fmt.Println("Саламалейкум")
	feature1.Feature1()

	for i := 0; i < 5; i++ {

		fmt.Println("Как же многа  строк")
	}
	feature2.Feature2()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	// if err := simple_sql.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Была создана таблица")
	// if err := simple_sql.InserRow(ctx, conn); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Задча успешно добавлена гугугага")
	if err := simple_sql.UpdateROW(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("Обновил запись в таблице успешно!")
}
