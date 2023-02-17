package main

import (
	"fmt"
	//  product "main/modules"
	"os"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	Products []Product
}

var inv Inventory

func main() {
	inventory := &Inventory{
		Products: []Product{
			{ID: 1, Name: "Product A", Price: 9.99, Quantity: 10},
			{ID: 2, Name: "Product B", Price: 19.99, Quantity: 5},
			{ID: 3, Name: "Product C", Price: 29.99, Quantity: 3},
		},
	}
	displayInventory(inventory)
	for {
		fmt.Println("Enter your command:")
		var command string
		fmt.Scanln(&command)
		switch command {
		case "add":
			var product Product
			inventory.add(product)
			displayInventory(inventory)
		case "remove":
			fmt.Println("Enter the ID of the product you want to remove: ")
			var id int
			fmt.Scan(&id)
			inventory.removeProduct(id)
			fmt.Println("Product removed.")
			displayInventory(inventory)
		case "display":
			displayInventory(inventory)
		case "update":
			fmt.Println("So you are updating Product!!")
			var product Product
			inventory.update(product)
			displayInventory(inventory)
		case "exit":
			os.Exit(0)
		}
	}

}

func (i *Inventory) update(p Product) {
	fmt.Print("Enter ID: ")
	fmt.Scan(&p.ID)
	fmt.Print("Enter name: ")
	fmt.Scan(&p.Name)
	fmt.Print("Enter price: ")
	fmt.Scan(&p.Price)
	fmt.Print("Enter quantity: ")
	fmt.Scan(&p.Quantity)
	if i.updateProduct(p) {
		fmt.Println("Product updated.")
	} else {
		fmt.Println("Product not found.")
	}

}

func (i *Inventory) updateProduct(p Product) bool {
	for index, product := range i.Products {
		if product.ID == p.ID {
			i.Products[index] = p
			return true
		}
	}
	return false
}

func (i *Inventory) add(p Product) {
	fmt.Print("Enter ID: ")
	fmt.Scan(&p.ID)
	fmt.Print("Enter name: ")
	fmt.Scan(&p.Name)
	fmt.Print("Enter price: ")
	fmt.Scan(&p.Price)
	fmt.Print("Enter quantity: ")
	fmt.Scan(&p.Quantity)
	i.addProduct(p)
	fmt.Println("Product added.")
}

func (i *Inventory) removeProduct(id int) bool {
	for index, product := range i.Products {
		if product.ID == id {
			i.Products = append(i.Products[:index], i.Products[index+1:]...)
			return true
		}
	}
	return false

}

func (i *Inventory) addProduct(p Product) {
	i.Products = append(i.Products, p)
}
func displayInventory(i *Inventory) {
	fmt.Println("Inventory:")
	for _, product := range i.Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Quantity: %d\n", product.ID, product.Name, product.Price, product.Quantity)
	}
	fmt.Println()
}
