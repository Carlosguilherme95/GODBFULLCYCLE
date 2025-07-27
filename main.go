package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main.go/database"
	"main.go/service"
)

func main() {
	db, err := database.DatabaseConection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("------CONECTADO AO BANCO DE DADOS--------")
	product := service.NewProduct("Iphone", 1500)
	err = service.InsertProduct(db, product)
	if err != nil {
		fmt.Println(err)
	}
	product.Price = 1100
	err = service.UpdateProduct(db, product)
	if err != nil {
		fmt.Println(err)
	}
	products, err := service.SelecAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Produto: %s - Preço: %.2f\n", p.Name, p.Price)
	}
	err = service.DeleteProduct(db, product.ID)
}
