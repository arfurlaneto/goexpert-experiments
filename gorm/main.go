package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		panic(err)
	}

	err = RunMigrations(db)
	if err != nil {
		panic(err)
	}

	DeleteAllProducts(db)
	InsertProducts(db)
	UpdateProducts(db)
	SelectProducts(db)

	err = UpdateWithLock(db)
	if err != nil {
		panic(err)
	}
}

func CreateConnection() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&Product{}, &Category{}, &Tag{}, &SerialNumber{})
}

func DeleteAllProducts(db *gorm.DB) {
	var serialNumbers []SerialNumber
	db.Unscoped().Find(&serialNumbers)
	if len(serialNumbers) > 0 {
		db.Unscoped().Delete(serialNumbers)
	}

	var products []Product
	db.Unscoped().Find(&products)
	if len(products) > 0 {
		db.Unscoped().Delete(products)
	}

	var tags []Tag
	db.Unscoped().Find(&tags)
	if len(tags) > 0 {
		db.Unscoped().Delete(tags)
	}

	var categories []Category
	db.Unscoped().Find(&categories)
	if len(categories) > 0 {
		db.Unscoped().Delete(categories)
	}
}

func InsertProducts(db *gorm.DB) {
	var offerTag Tag = Tag{Name: "offer"}
	db.Create(&offerTag)

	var peripheralTag Tag = Tag{Name: "peripheral"}
	db.Create(&peripheralTag)

	var newTag = Tag{Name: "new"}
	db.Create(&newTag)

	var category Category
	var product Product

	category = Category{Name: "Technology"}
	db.Create(&category)

	product = Product{Name: "Notebook", Price: 3000.0, CategoryID: category.ID, Tags: []Tag{newTag, offerTag}}
	db.Create(&product)
	db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: product.ID})

	product = Product{Name: "Monitor", Price: 700.0, CategoryID: category.ID, Tags: []Tag{newTag}}
	db.Create(&product)
	db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: product.ID})

	product = Product{Name: "Keyboard", Price: 50.0, CategoryID: category.ID, Tags: []Tag{peripheralTag}}
	db.Create(&product)
	db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: product.ID})

	product = Product{Name: "Mouse", Price: 20.0, CategoryID: category.ID, Tags: []Tag{peripheralTag}}
	db.Create(&product)
	db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: product.ID})
}

func UpdateProducts(db *gorm.DB) {
	var product Product
	db.First(&product)

	product.Name = product.Name + " (UPDATED)"
	db.Save(product)
}

func SelectProducts(db *gorm.DB) {
	var product Product
	var products []Product

	db.Preload("Category").Preload("SerialNumber").Preload("Tags").First(&product)
	printProduct(">> First", &product)

	db.Preload("Category").Preload("SerialNumber").Preload("Tags").First(&product, product.ID)
	printProduct(fmt.Sprintf(">> ID = %d", product.ID), &product)

	db.Preload("Category").Preload("SerialNumber").Preload("Tags").First(&product, "name like ?", "%Notebook%")
	printProduct(">> name LIKE %Notebook%", &product)

	db.Preload("Category").Preload("SerialNumber").Preload("Tags").Where("price > ?", 20).Find(&products)
	printProducts(">> price > 20", products)

	db.Preload("Category").Preload("SerialNumber").Preload("Tags").Find(&products)
	printProducts(">> All", products)

	var categories []Category
	db.Model(&Category{}).
		Preload("Products.Category").
		Preload("Products.SerialNumber").
		Preload("Products.Tags").
		Find(&categories)

	for _, category := range categories {
		printProducts(fmt.Sprintf("CATEGORY \"%s\"", category.Name), products)
	}
}

func UpdateWithLock(db *gorm.DB) error {
	tx := db.Begin()

	var c Category
	err := tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c).Error
	if err != nil {
		return err
	}

	c.Name = "Eletronics"
	tx.Debug().Save(&c)

	tx.Commit()
	return nil
}
