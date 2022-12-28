package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type Widget struct {
	ID          int `gorm:"AUTO_INCREMENT;primary_key"`
	Name        string
	Description *string
}

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

func CreateRandomWidget() *Widget {
	widget := &Widget{
		Name:        fmt.Sprintf("Widget: %s", strconv.Itoa(random.Intn(1000))),
	}

	return widget
}
