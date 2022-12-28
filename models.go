package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Widget struct {
	ID          uint `gorm:"AUTO_INCREMENT;primary_key"`
	Name        string
	Description *string
	Inventory   []Inventory
}

type Location struct {
	ID   uint `gorm:"AUTO_INCREMENT;PRIMARY_KEY`
	Name string
}

type Inventory struct {
	ID         uint `gorm:AUTO_INCREMENT;PRIMARY_KEY`
	LocationID uint
	WidgetID   uint
	Quantity   uint
}

func (_ Inventory) TableName() string {
	return "inventory"
}

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

func CreateRandomWidget() *Widget {
	widget := &Widget{
		Name: fmt.Sprintf("Widget: %s", strconv.Itoa(random.Intn(1000))),
	}

	return widget
}
