// 代码生成时间: 2025-10-20 00:38:55
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "log"
)

// Product represents a product in the supply chain
type Product struct {
    ID        uint   "db:id"
    Name      string "db:name"
    Manufacturer string   "db:manufacturer"
    ProductionDate string   "db:production_date"
    SupplierID uint   "db:supplier_id"
}

// ProductResource represents the resource for products
type ProductResource struct {
    ID        uint   "json:"id"`
    Name      string "json:"name"`
    Manufacturer string   "json:"manufacturer"`
    ProductionDate string   "json:"production_date"`
    SupplierID uint   "json:"supplier_id"`
}

// ProductService handles business logic for products
type ProductService struct {
    DB *pop.Connection
}

// NewProductService creates a new ProductService with a database connection
func NewProductService(db *pop.Connection) *ProductService {
    return &ProductService{DB: db}
}

// AddProduct adds a new product to the supply chain
func (s *ProductService) AddProduct(name, manufacturer, productionDate string, supplierID uint) (*Product, error) {
    product := Product{
        Name:      name,
        Manufacturer: manufacturer,
        ProductionDate: productionDate,
        SupplierID: supplierID,
    }
    err := s.DB.Create(&product)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

// GetProduct retrieves a product by ID
func (s *ProductService) GetProduct(id uint) (*Product, error) {
    var product Product
    err := s.DB.FindBy("id", id, &product)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

// Main function to run the Buffalo application
func main() {
    app := buffalo.Automatic(buffalo.Options{
        AppName: "supply-chain-traceability",
        NoBanner: true,
    })

    // Register the ProductService
    app.Use(func(next buffalo.Handler) buffalo.Handler {
        return func(c buffalo.Context) error {
            tx, err := app.DB.Begin()
            if err != nil {
                return err
            }
            defer tx.Rollback()

            c.Set("DB", tx)
            return next(c)
        }
    })

    // Define routes for products
    app.Resource("/products", NewProductResource())

    // Run the application
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}

// ProductResource is a Buffalo resource for handling HTTP requests
type ProductResource struct{}

// List handles GET requests to /products
func (p *ProductResource) List(c buffalo.Context) error {
    tx := c.Value("DB").(*pop.Connection)
    var products []Product
    if err := tx.All(&products); err != nil {
        return c.Error(500, err)
    }
    return c.Render(200, r.JSON(products))
}

// Show handles GET requests to /products/{id}
func (p *ProductResource) Show(c buffalo.Context) error {
    id := c.Param("id")
    tx := c.Value("DB").(*pop.Connection)
    var product Product
    if err := tx.Find(&product, id); err != nil {
        return c.Error(404, err)
    }
    return c.Render(200, r.JSON(product))
}

// Create handles POST requests to /products
func (p *ProductResource) Create(c buffalo.Context) error {
    var product ProductResource
    if err := c.Bind(&product); err != nil {
        return c.Error(400, err)
    }
    service := NewProductService(c.Value("DB").(*pop.Connection))
    dbProduct, err := service.AddProduct(product.Name, product.Manufacturer, product.ProductionDate, product.SupplierID)
    if err != nil {
        return c.Error(500, err)
    }
    return c.Render(201, r.JSON(dbProduct))
}

// Update handles PUT requests to /products/{id}
func (p *ProductResource) Update(c buffalo.Context) error {
    id := c.Param("id")
    var product ProductResource
    if err := c.Bind(&product); err != nil {
        return c.Error(400, err)
    }
    tx := c.Value("DB").(*pop.Connection)
    var dbProduct Product
    if err := tx.Find(&dbProduct, id); err != nil {
        return c.Error(404, err)
    }
    dbProduct.Name = product.Name
    dbProduct.Manufacturer = product.Manufacturer
    dbProduct.ProductionDate = product.ProductionDate
    dbProduct.SupplierID = product.SupplierID
    if err := tx.Update(&dbProduct); err != nil {
        return c.Error(500, err)
    }
    return c.Render(200, r.JSON(dbProduct))
}

// Destroy handles DELETE requests to /products/{id}
func (p *ProductResource) Destroy(c buffalo.Context) error {
    id := c.Param("id")
    tx := c.Value("DB").(*pop.Connection)
    var product Product
    if err := tx.Find(&product, id); err != nil {
        return c.Error(404, err)
    }
    if err := tx.Destroy(&product); err != nil {
        return c.Error(500, err)
    }
    return c.Render(200, r.JSON(genericMap{
        "id": id,
        "message": "Product successfully deleted",
    }))
}
