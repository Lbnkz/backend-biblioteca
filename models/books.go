package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID            int       `gorm:"column:id;primary_key"` 
	Titulo        string    `gorm:"column:titulo"`
	Autor         string    `gorm:"column:autor"`
	AnoPublicacao int       `gorm:"column:anopublicacao"`
}
