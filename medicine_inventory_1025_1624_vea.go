// 代码生成时间: 2025-10-25 16:24:20
package main

import (
    "buffalo"
    "buffalo/buffalo-plugins"
    "github.com/markbates/pkger"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Medicine represents the medicine in the inventory.
type Medicine struct {
    gorm.Model
    Name        string `json:"name"`
    Quantity    int    `json:"quantity"`
    Description string `json:"description"`
}

// MedicineInventory is the main application.
type MedicineInventory struct{
    *buffalo.App
}

// DB is the application-wide database connection.
var DB, _ = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

// CreateMedicine is the handler that will add a new medicine to the inventory.
func (app *MedicineInventory) CreateMedicine(c buffalo.Context) error {
    var medicine Medicine
    if err := c.Bind(&medicine); err != nil {
        return err
    }
    if err := DB.Create(&medicine).Error; err != nil {
        return err
    }
    return c.Render(201, r.Data(medicine))
}

// UpdateMedicine is the handler that will update an existing medicine in the inventory.
func (app *MedicineInventory) UpdateMedicine(c buffalo.Context) error {
    medicineID := c.Param("id")
    var medicine Medicine
    if err := DB.First(&medicine, medicineID).Error; err != nil {
        return c.Error(404, err)
    }
    if err := c.Bind(&medicine); err != nil {
        return err
    }
    if err := DB.Save(&medicine).Error; err != nil {
        return err
    }
    return c.Render(200, r.Data(medicine))
}

// DeleteMedicine is the handler that will remove a medicine from the inventory.
func (app *MedicineInventory) DeleteMedicine(c buffalo.Context) error {
    medicineID := c.Param("id")
    var medicine Medicine
    if err := DB.Delete(&medicine, medicineID).Error; err != nil {
        return err
    }
    return c.Render(200, r.Data(medicine))
}

// GetMedicine is the handler that will retrieve a medicine by ID.
func (app *MedicineInventory) GetMedicine(c buffalo.Context) error {
    medicineID := c.Param("id")
    var medicine Medicine
    if err := DB.First(&medicine, medicineID).Error; err != nil {
        return c.Error(404, err)
    }
    return c.Render(200, r.Data(medicine))
}

// GetMedicines is the handler that will retrieve all medicines in the inventory.
func (app *MedicineInventory) GetMedicines(c buffalo.Context) error {
    var medicines []Medicine
    if err := DB.Find(&medicines).Error; err != nil {
        return c.Error(404, err)
    }
    return c.Render(200, r.Data(medicines))
}

// main is the entry point for the application.
func main() {
    app := buffalo.Automatic(buffalo.Pluralize("medicine"), "{root}/templates")
    app.GET("/", HomeHandler)
    app.POST("/medicines", func(c buffalo.Context) error { return c.Render(200, r.String("Welcome to the Medicine Inventory System!")) })
    app.Resource("/medicines", &MedicineInventory{})
    plugins.Start(app)
}

// HomeHandler is a simple handler that prints a welcome message.
func HomeHandler(c buffalo.Context) error {
    return c.Render(200, r.String("Welcome to the Medicine Inventory System!"))
}
