package service

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"main.go/entities"
)

// retorna um ponteiro para a minha Struct (local na memória *entities.Product)
// função para criar novo produto
func NewProduct(name string, price float64) *entities.Product {
	return &entities.Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
func InsertProduct(db *sql.DB, product *entities.Product) error {
	stmt, err := db.Prepare("INSERT INTO Product (ID ,Name, Price) VALUES (?,?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	fmt.Println("------ PRODUTO INSERIDO COM SUCESSO --------")
	return nil
}
func UpdateProduct(db *sql.DB, product *entities.Product) error {
	smtp, err := db.Prepare("UPDATE Product SET Name = ?, Price = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer smtp.Close()
	_, err = smtp.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	fmt.Println("------ PRODUTO ATUALZIADO SOM SUCESSO -----")
	return nil
}
func SelectOneProduct(db *sql.DB, id string) (*entities.Product, error) {
	var p entities.Product
	// query row significa que está buscando apenas uma linha
	err := db.QueryRow("SELECT * FROM Product WHERE ID = ?", id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Product: %v, possui o preço de R$ %.2f", p.Name, p.Price)
	return &p, nil
}

// retorna um slice dos produtos (slice é como se fosse um array)
func SelecAllProducts(db *sql.DB) ([]entities.Product, error) {
	rows, err := db.Query("SELECT ID, Name, Price FROM Product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []entities.Product
	// faço uma lista de product e depois jogo eles no meu slice, como se fosse um array.push
	for rows.Next() {
		var p entities.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
func DeleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM Product WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("------ PRODUTO DELETADO COM SUCESSO --------")
	return nil
}
