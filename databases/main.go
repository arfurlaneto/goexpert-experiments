package main

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price int
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	newId := uuid.New().String()

	_, err = stmt.Exec(newId, "New Product", 1000)
	if err != nil {
		panic(err)
	}

	stmt, err = db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(newId).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	println(string(jsonData))

	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}

	jsonData, err = json.Marshal(products)
	if err != nil {
		panic(err)
	}

	println(string(jsonData))

	stmt, err = db.Prepare("delete from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(newId)
	if err != nil {
		panic(err)
	}
}
