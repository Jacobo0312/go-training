package structs

// Global Slice for Products
var products []Product

//Interface

type ProductInterface interface {
	Price() float64
}

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

//Functions

// GetAll
func GetAll() []Product {
	return products
}

// FindByID
func FindByID(id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

// Product Small
type ProductSmall struct {
	BaseProduct Product
}

// Methods
func (p ProductSmall) Price() float64 {
	return p.BaseProduct.Price
}

// Product Medium
type ProductMedium struct {
	BaseProduct Product
}

//Methods

func (p ProductMedium) Price() float64 {
	return p.BaseProduct.Price * 1.03
}

// Product Large
type ProductLarge struct {
	BaseProduct Product
}

// Methods
func (p ProductLarge) Price() float64 {
	return p.BaseProduct.Price*1.06 + 2500
}

// Factory
func CreateProduct(p Product, typeProduct string) ProductInterface {
	switch typeProduct {
	case "small":
		return &ProductSmall{BaseProduct: p}
	case "medium":
		return &ProductMedium{BaseProduct: p}
	case "large":
		return &ProductLarge{BaseProduct: p}
	default:
		return &ProductSmall{BaseProduct: p}
	}
}
