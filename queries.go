package main

import (
	"github.com/jinzhu/gorm"
)

func ReadAllWidgets(db *gorm.DB) ([]Widget, error) {
	widgets := make([]Widget, 0)
	err := db.Find(&widgets).Error

	if err != nil {
		return nil, err
	}

	return widgets, nil
}

func ReadInventoryFor(db *gorm.DB, widget Widget) ([]Inventory, error) {
	inventory := make([]Inventory, 0)
	err := db.Debug().Model(widget).Association("Inventory").Find(&inventory).Error

	if err != nil {
		return nil, err
	}

	return inventory, nil
}
