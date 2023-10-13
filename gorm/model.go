package main

import "gorm.io/gorm"

type Product struct {
	ID           int `goorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category     // many to one
	Tags         []Tag        `gorm:"many2many:products_tags;constraint:OnDelete:CASCADE"` // many to many
	SerialNumber SerialNumber // one to one
	gorm.Model
}

type Category struct {
	ID       int `goorm:"primaryKey"`
	Name     string
	Products []Product // one to many
	gorm.Model
}

type Tag struct {
	ID       int `goorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_tags;constraint:OnDelete:CASCADE"` // many to many
	gorm.Model
}

type SerialNumber struct {
	ID        int `goorm:"primaryKey"`
	Number    string
	ProductID int
	gorm.Model
}
