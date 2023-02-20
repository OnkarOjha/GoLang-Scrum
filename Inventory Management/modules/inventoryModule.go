package product

type Product struct {
	ID int
	Name string 
	Price float64
	Quantity int
}

type Inventory struct{
	Products []Product
}