// 代码生成时间: 2025-10-18 11:38:07
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/generators"
    "github.com/markbates/inflect"
    "github.com/markbates/validate"
)

// Metadata is the struct representing the metadata model
type Metadata struct {
    ID        uint   `db:"id" json:"id"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
    DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
    // Add other fields as needed
    // Name         string   `db:"name" json:"name"`
}

// Validate is a model method that performs validation
func (m *Metadata) Validate(tx *pop.Connection, isCreate bool) error {
    // Implement validation logic if necessary
    return nil
}

func main() {
    // Initialize the Buffalo app
    app := buffalo.Automatic()

    // Define the routes for the metadata management
    app.GET("/metadata", metadataList)
    app.POST("/metadata", metadataCreate)
    app.GET("/metadata/{param_id}", metadataShow)
    app.PUT("/metadata/{param_id}", metadataUpdate)
    app.DELETE("/metadata/{param_id}", metadataDelete)

    // Start the Buffalo server
    app.Serve()
}

// metadataList handles GET requests for listing metadata
func metadataList(c buffalo.Context) error {
    var metadata []Metadata
    // Use the database connection to retrieve the metadata
    if err := c.Value("db").(*pop.Connection).Where("deleted_at IS NULL").All(&metadata); err != nil {
        return err
    }
    return c.Render(200, r.JSON(metadata))
}

// metadataCreate handles POST requests for creating metadata
func metadataCreate(c buffalo.Context) error {
    var m Metadata
    // Bind the request body to the metadata struct
    if err := c.Bind(&m); err != nil {
        return err
    }
    // Validate the metadata
    if err := m.Validate(c.Value("db").(*pop.Connection), true); err != nil {
        return err
    }
    // Save the new metadata
    if err := c.Value("db").(*pop.Connection).Create(&m); err != nil {
        return err
    }
    return c.Render(201, r.JSON(m))
}

// metadataShow handles GET requests for showing a single metadata
func metadataShow(c buffalo.Context) error {
    id := c.Param("param_id")
    // Retrieve the metadata by ID
    var m Metadata
    if err := c.Value("db").(*pop.Connection).Find(&m, id); err != nil {
        return err
    }
    return c.Render(200, r.JSON(m))
}

// metadataUpdate handles PUT requests for updating metadata
func metadataUpdate(c buffalo.Context) error {
    id := c.Param("param_id\)
    var m Metadata
    if err := c.Value("db").(*pop.Connection).Find(&m, id); err != nil {
        return err
    }
    // Bind the request body to the metadata struct
    if err := c.Bind(&m); err != nil {
        return err
    }
    // Validate the metadata
    if err := m.Validate(c.Value("db").(*pop.Connection), false); err != nil {
        return err
    }
    // Update the metadata
    if err := c.Value("db").(*pop.Connection).Update(&m); err != nil {
        return err
    }
    return c.Render(200, r.JSON(m))
}

// metadataDelete handles DELETE requests for deleting metadata
func metadataDelete(c buffalo.Context) error {
    id := c.Param("param_id\)
    var m Metadata
    if err := c.Value("db").(*pop.Connection).Find(&m, id); err != nil {
        return err
    }
    // Soft delete the metadata
    m.Delete()
    return c.Render(200, r.JSON(m))
}
