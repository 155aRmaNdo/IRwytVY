// 代码生成时间: 2025-10-30 02:52:51
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "github.com/gobuffalo/envy"
    "log"
)

// DeviceFirmwareUpdate represents a firmware update for a device
type DeviceFirmwareUpdate struct {
    // ID is the unique identifier for the firmware update
    ID uint `db:""`
    // DeviceID is the identifier for the device being updated
    DeviceID uint `db:"index"`
    // FirmwareVersion is the version of the firmware being updated
    FirmwareVersion string `db:"size:255"`
    // Status indicates the current status of the update
    Status string `db:"size:255"`
    // CreatedAt is the timestamp when the update was created
    CreatedAt string `db:""`
    // UpdatedAt is the timestamp when the update was last updated
    UpdatedAt string `db:""`
}

// CreateFirmwareUpdate is a handler function that creates a new firmware update
func CreateFirmwareUpdate(c buffalo.Context) error {
    // Decode the JSON payload into a DeviceFirmwareUpdate struct
    var update DeviceFirmwareUpdate
    if err := c.Bind(&update); err != nil {
        return err
    }

    // Validate the update data
    if update.FirmwareVersion == "" {
        return buffalo.NewError("Firmware version is required")
    }

    // Save the new firmware update to the database
    if err := update.Save(c); err != nil {
        return err
    }

    // Return a success response
    return c.Render(200, r.JSON(update))
}

// UpdateFirmwareStatus is a handler function that updates the status of a firmware update
func UpdateFirmwareStatus(c buffalo.Context) error {
    // Get the ID of the firmware update from the URL parameter
    id, err := envy.ParseUint(c.Param("id"))
    if err != nil {
        return err
    }

    // Find the firmware update in the database
    var update DeviceFirmwareUpdate
    if err := update.FindByID(id, c); err != nil {
        return err
    }

    // Decode the JSON payload into the update struct
    if err := c.Bind(&update); err != nil {
        return err
    }

    // Save the updated status to the database
    if err := update.Save(c); err != nil {
        return err
    }

    // Return a success response
    return c.Render(200, r.JSON(update))
}

// main is the entry point for the Buffalo application
func main() {
    // Create a new Buffalo application
    app := buffalo.Automatic()

    // Add middleware to parse JSON requests
    app.Use((buffalo.MiddlewareFunc)(func(c buffalo.Context, next buffalo.Handler) error {
        return next(c)
    }))

    // Add the CreateFirmwareUpdate handler to the /firmware-updates route
    app.POST("/firmware-updates", CreateFirmwareUpdate)

    // Add the UpdateFirmwareStatus handler to the /firmware-updates/{id}/status route
    app.PATCH("/firmware-updates/{id}/status", UpdateFirmwareStatus)

    // Start the Buffalo application
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}
