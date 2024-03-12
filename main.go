package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testTask/db"
	"testTask/repository"
)

func main() {
	dbMySql, err := sqlx.Connect("mysql", "user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return
	}
	defer dbMySql.Close()

	err = db.CreateTable(dbMySql)

	if err != nil {
		panic("error create table")
	}

	_ = repository.NewCitiesRepository(dbMySql)

	fmt.Println("Таблица успешно создана!")
}
