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

	err = db.Create(widget).Error
	if err != nil {
		panic(err)
	}

	err = db.Create(&Inventory{
		Quantity:   1,
		LocationID: 1,
		WidgetID:   widget.ID,
	}).Error
	if err != nil {
		panic(err)
	}

	widgets, err := ReadAllWidgets(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("All widgets")
	printWidgets(widgets)

	inventoryLines, err := ReadInventoryFor(db, *widget)
	if err != nil {
		panic(err)
	}

	fmt.Println("All inventory for the new widget")
	for _, inventory := range inventoryLines {
		fmt.Printf("%d %s; quantity %d\n", widget.ID, widget.Name, inventory.Quantity)
	}
}
