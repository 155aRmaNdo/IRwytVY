// 代码生成时间: 2025-10-12 18:33:37
package main

import (
    "log"
    "net/http"
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/gobuffalo/x/xhttp"
    "your_application/models" // Replace with your actual import path
)

// 定义医疗资源调度的模型结构
type MedicalResource struct {
    ID       uint   "db:"id"
    Name     string "db:"name"
    Quantity int    "db:"quantity"
    pop.Model
}

// NewMedicalResourceHandler 处理医疗资源的创建请求
func NewMedicalResourceHandler(db *pop.Connection) buffalo.Handler {
    return func(c buffalo.Context) error {
        // 解析请求体中的JSON数据
        var res MedicalResource
        if err := c.Bind(&res); err != nil {
            return errors.New("Failed to bind request data")
        }

        // 将医疗资源添加到数据库
        if err := db.Create(&res); err != nil {
            return errors.New("Failed to create medical resource")
        }

        // 返回创建的医疗资源
        return c.Render(200, r.JSON(res))
    }
}

// EditMedicalResourceHandler 处理医疗资源的更新请求
func EditMedicalResourceHandler(db *pop.Connection) buffalo.Handler {
    return func(c buffalo.Context) error {
        // 解析ID并查询医疗资源
        id := c.Param("id")
        var res MedicalResource
        if err := db.Find(&res, id); err != nil {
            return errors.New("Failed to find medical resource")
        }

        // 解析请求体中的JSON数据
        if err := c.Bind(&res); err != nil {
            return errors.New("Failed to bind request data\)
        }

        // 更新医疗资源信息
        if err := db.Update(&res); err != nil {
            return errors.New("Failed to update medical resource\)
        }

        // 返回更新的医疗资源
        return c.Render(200, r.JSON(res))
    }
}

func main() {
    // 创建Buffalo应用
    app := buffalo.Automatic()

    // 添加路由
    app.Resource("/medical-resources", MedicalResource{},
        buffalo Routes{
            {
                "path": "/medical-resources",
                "method": "GET",
                "action": MedicalResourceListHandler(db),
            },
            {
                "path": "/medical-resources",
                "method": "POST",
                "action": NewMedicalResourceHandler(db),
            },
            {
                "path": "/medical-resources/{id}",
                "method": "PUT",
                "action": EditMedicalResourceHandler(db),
            },
            {
                "path": "/medical-resources/{id}",
                "method": "DELETE",
                "action": DeleteMedicalResourceHandler(db),
            },
        },
    )

    // 添加中间件
    app.Use(middleware.ParameterLogger)
    app.Use(middleware.RecoveryHandler)
    app.Use(xhttp.MethodOverrideHandler)

    // 设置数据库连接
    app dbs buffalo.DB
    app = app.WithDB(dbs)

    // 启动Buffalo应用
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}
