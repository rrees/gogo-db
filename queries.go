package main

import (
	"github.com/jinzhu/gorm"
)

func ReadAllWidgets(db *gorm.DB) []Widget {
  widgets := make([]Widget, 0)
	db.Find(&widgets)

  return widgets
}
