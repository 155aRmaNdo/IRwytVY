// 代码生成时间: 2025-10-07 21:11:39
package main

import (
    "buffalo(buffalo.dev)"
    "github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
    "golang.org/x/crypto/bcrypt"
    "log"
)

// AuthService 负责处理用户身份认证
type AuthService struct {
    DB *popmw.Connection
}

// NewAuthService 创建一个新的AuthService实例
func NewAuthService(db *popmw.Connection) *AuthService {
    return &AuthService{DB: db}
}

// Authenticate 用户登录认证
func (a *AuthService) Authenticate(email, password string) (bool, error) {
    // 1. 根据邮箱查询用户
    var user User
    err := a.DB.Where("email = ?", email).First(&user)
    if err != nil {
        log.Println("Error finding user: ", err)
        return false, err
    }

    // 2. 验证密码
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        log.Println("Error comparing password: ", err)
        return false, err
    }

    return true, nil
}

// User 定义用户结构体
type User struct {
    ID       uint
    Email    string
    Password string `db:"password"`
}
