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
