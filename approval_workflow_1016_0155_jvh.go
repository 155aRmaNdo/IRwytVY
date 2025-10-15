// 代码生成时间: 2025-10-16 01:55:29
package main

import (
    "log"
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware/csrf"
    "github.com/gobuffalo/buffalo/middleware/i18n"
    "github.com/gobuffalo/buffalo/middleware/ssl"
    "github.com/gobuffalo/envy"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

// 初始化数据库连接
func init() {
    var err error
    db, err = gorm.Open(sqlite.Open(envy.Get("DB_NAME", "approval_workflow.db")), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: %v", err)
    }
    
    // 自动迁移模式
    db.AutoMigrate(&Approval{}, &Process{}, &Step{}, &ApprovalStep{})
}

// Approval 审批对象
type Approval struct {
    gorm.Model
    Name string
    Description string
    Process Process `gorm:"foreignKey:ApprovalID"`
}

// Process 审批流程
type Process struct {
    gorm.Model
    Name string
    Description string
    Approval Approval `gorm:"foreignKey:ProcessID"`
    Steps []Step `gorm:"foreignKey:ProcessID"`
}

// Step 审批步骤
type Step struct {
    gorm.Model
    Name string
    Description string
    ProcessID uint
    NextStepID *uint
}

// ApprovalStep 审批步骤实例
type ApprovalStep struct {
    gorm.Model
    ApprovalID uint
    StepID uint
    Completed bool
    Comments string
}

// Models 定义模型
var Models = []interface{}{
    &Approval{},
    &Process{},
    &Step{},
    &ApprovalStep{},
}

// app 应用实例
var app *buffalo.App
var db *gorm.DB

func main() {
    app = buffalo.Automatic(buffalo.Options{
        Env: envy.Get("GO_ENV", "development"),
        PrettyErrors: envy.GetBool("PRETTY_ERRORS", true),
    })
    
    // 设置中间件
    app.Use(csrf.New)
    app.Use(i18n.New)
    app.Use(ssl.ForceSSL{"except": []string{"/", "/login", "/api"}}))
    
    // 定义路由
    app.GET("/", HomeHandler)
    app.GET("/processes", ProcessListHandler)
    app.POST("/processes", ProcessCreateHandler)
    app.PATCH("/processes/{pid}", ProcessUpdateHandler)
    
    // 启动应用
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}

// HomeHandler 首页处理函数
func HomeHandler(c buffalo.Context) error {
    return c.Render(200, r.HTML("index.html"))
}

// ProcessListHandler 审批流程列表处理函数
func ProcessListHandler(c buffalo.Context) error {
    var processes []Process
    if err := db.Find(&processes).Error; err != nil {
        return c.Error(500, err)
    }
    return c.Render(200, r.HTML{"processes.html", processes})
}

// ProcessCreateHandler 审批流程创建处理函数
func ProcessCreateHandler(c buffalo.Context) error {
    var process Process
    if err := c.Bind(&process); err != nil {
        return c.Error(400, err)
    }
    if err := db.Create(&process).Error; err != nil {
        return c.Error(500, err)
    }
    return c.Redirect(302, "/processes")
}

// ProcessUpdateHandler 审批流程更新处理函数
func ProcessUpdateHandler(c buffalo.Context) error {
    var process Process
    if err := c.Bind(&process); err != nil {
        return c.Error(400, err)
    }
    if err := db.Model(&process).Updates(process).Error; err != nil {
        return c.Error(500, err)
    }
    return c.Redirect(302, "/processes")
}