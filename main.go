package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func printWidgets(widgets []Widget) {
	for n, widget := range widgets {
		fmt.Printf("%d %s\n", n, widget.Name)
	}
}

func main() {
	fmt.Println("Hello world")

	dbString := os.Getenv("DB_STRING")

	if dbString == "" {
		panic("The connection string DB_STRING is not defined")
	}

	db, err := gorm.Open("mysql", dbString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	defer db.Close()

	widget := CreateRandomWidget()

	db.Create(widget)

	printWidgets(ReadAllWidgets(db))
}
