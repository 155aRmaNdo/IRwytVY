// 代码生成时间: 2025-10-28 18:51:04
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/meta/inflect"
    "github.com/markbates/pkger"
    "log"
)

// ClinicalTrial 定义临床试验模型
type ClinicalTrial struct {
    ID        uint   `db:"id"`
    Title     string `db:"title"`
    StartDate string `db:"start_date"`
    EndDate   string `db:"end_date"`
    Status    string `db:"status"`
}

// ClinicalTrialResource 定义资源
type ClinicalTrialResource struct {
    // 可以添加更多属性和方法
}

// NewClinicalTrialResource 初始化资源
func NewClinicalTrialResource() *ClinicalTrialResource {
    return &ClinicalTrialResource{}
}

// List 列出所有临床试验
func (resource *ClinicalTrialResource) List(c buffalo.Context) error {
    // 获取模型
    var clinicalTrials []ClinicalTrial
    if err := c.DB().Query("SELECT * FROM clinical_trials").Bind(&clinicalTrials); err != nil {
        // 错误处理
        return buffalo.NewError("An error occurred while retrieving clinical trials").SetType(buffalo.StatusError)
    }
    // 返回结果
    return c.Render(200, r.JSON(clinicalTrials))
}

// Show 查看单个临床试验
func (resource *ClinicalTrialResource) Show(c buffalo.Context) error {
    // 获取ID
    id := c.Param("id")
    // 根据ID查询临床试验
    var clinicalTrial ClinicalTrial
    if err := c.DB().Where("id = ?", id).First(&clinicalTrial); err != nil {
        // 错误处理
        return buffalo.NewError("Clinical trial not found").SetType(buffalo.StatusError)
    }
    // 返回结果
    return c.Render(200, r.JSON(clinicalTrial))
}

// Create 创建新临床试验
func (resource *ClinicalTrialResource) Create(c buffalo.Context) error {
    // 解析请求体
    var clinicalTrial ClinicalTrial
    if err := c.Bind(&clinicalTrial); err != nil {
        return err
    }
    // 插入数据库
    if err := c.DB().Create(&clinicalTrial); err != nil {
        // 错误处理
        return buffalo.NewError("An error occurred while creating clinical trial").SetType(buffalo.StatusError)
    }
    // 返回结果
    return c.Render(201, r.JSON(clinicalTrial))
}

// Update 更新临床试验信息
func (resource *ClinicalTrialResource) Update(c buffalo.Context) error {
    // 获取ID
    id := c.Param("id")
    // 查询临床试验
    var clinicalTrial ClinicalTrial
    if err := c.DB().Where("id = ?", id).First(&clinicalTrial); err != nil {
        return buffalo.NewError("Clinical trial not found").SetType(buffalo.StatusError)
    }
    // 解析请求体
    if err := c.Bind(&clinicalTrial); err != nil {
        return err
    }
    // 更新数据库
    if err := c.DB().Save(&clinicalTrial); err != nil {
        return buffalo.NewError("An error occurred while updating clinical trial").SetType(buffalo.StatusError)
    }
    // 返回结果
    return c.Render(200, r.JSON(clinicalTrial))
}

// Delete 删除临床试验
func (resource *ClinicalTrialResource) Delete(c buffalo.Context) error {
    // 获取ID
    id := c.Param("id")
    // 查询临床试验
    var clinicalTrial ClinicalTrial
    if err := c.DB().Where("id = ?", id).First(&clinicalTrial); err != nil {
        return buffalo.NewError("Clinical trial not found").SetType(buffalo.StatusError)
    }
    // 删除数据库记录
    if err := c.DB().Destroy(&clinicalTrial); err != nil {
        return buffalo.NewError("An error occurred while deleting clinical trial").SetType(buffalo.StatusError)
    }
    // 返回结果
    return c.Render(204, r.NoContent())
}

func main() {
    app := buffalo.Automatic()
    // 路由
    app.GET("/clinical-trials", NewClinicalTrialResource().List)
    app.GET("/clinical-trials/{id}", NewClinicalTrialResource().Show)
    app.POST("/clinical-trials", NewClinicalTrialResource().Create)
    app.PUT("/clinical-trials/{id}", NewClinicalTrialResource().Update)
    app.DELETE("/clinical-trials/{id}", NewClinicalTrialResource().Delete)

    // 应用运行
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }
}