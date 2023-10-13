package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func printProduct(title string, product *Product) {
	printProducts(title, []Product{*product})
}

func printProducts(title string, products []Product) {
	fmt.Println(title)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "S/N", "NAME", "PRICE", "CATEGORY", "TAGS"})

	for _, product := range products {
		tagsOutput := ""
		for _, tag := range product.Tags {
			tagsOutput = fmt.Sprintf("%s #%s", tagsOutput, tag.Name)
		}

		t.AppendRows([]table.Row{{
			product.ID,
			product.SerialNumber.Number,
			fmt.Sprintf("%-20s", product.Name),
			product.Price,
			product.Category.Name,
			fmt.Sprintf("%-20s", tagsOutput),
		}})
		t.AppendSeparator()
	}

	t.Render()
	fmt.Println()
}
