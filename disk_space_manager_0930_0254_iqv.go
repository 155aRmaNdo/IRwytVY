// 代码生成时间: 2025-09-30 02:54:23
package main

import (
    "buffalo"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "syscall"
    "unsafe"
)

// DiskSpaceManager defines a struct to handle disk space operations
type DiskSpaceManager struct {
    Path string
}

// NewDiskSpaceManager creates a new instance of DiskSpaceManager
func NewDiskSpaceManager(path string) *DiskSpaceManager {
    return &DiskSpaceManager{Path: path}
}

// GetDiskUsage returns the disk usage statistics for the given path
func (dsm *DiskSpaceManager) GetDiskUsage() (uint64, error) {
    // Get the file system statistics
    var stat syscall.Statfs_t
    if err := syscall.Statfs(dsm.Path, &stat); err != nil {
        return 0, err
    }

    // Calculate the total and free space in bytes
    totalSpace := stat.Blocks * uint64(stat.Bsize)
    freeSpace := stat.Bfree * uint64(stat.Bsize)

    return totalSpace, nil
}

func main() {
    // Initialize Buffalo application
    app := buffalo.New(buffalo.Options{})

    // Define a route for disk space check
    app.GET("/disk", func(c buffalo.Context) error {
        path := c.Param("path")
        if path == "" {
            path = "." // Default to current directory if no path is provided
        }

        dsm := NewDiskSpaceManager(path)
        totalSpace, err := dsm.GetDiskUsage()
        if err != nil {
            log.Printf("Error getting disk usage: %s", err)
            return c.Render(500, r.String("Error checking disk space"))
        }

        // Render the disk usage as JSON
        return c.Render(200, r.JSON(map[string]interface{}{
            "totalSpace": totalSpace,
        }))
    })

    // Start the Buffalo application
    app.Serve()
}
