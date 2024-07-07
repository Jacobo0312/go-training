package structs

// Global Slice for Products
var products []Product

// Product struct
// ID, Name, Price, Description y Category.
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

//Methods

// Save
func (p *Product) Save() {
	products = append(products, *p)
}

// GetAll
func (p *Product) GetAll() []Product {
	return products
}

// FindByID
func (p *Product) FindByID(id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
