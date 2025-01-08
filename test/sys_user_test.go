package test

import (
	"go-admin/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	models.NewGormDB()
	db := models.DB
	user := models.SysUser{
		UserName: "admin",
		PassWord: "admin123",
		Phone:    "18666666666",
		Avatar:   "avatar.png",
	}
	err := models.CreateSysUser(db, &user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// 从数据库中检索用户
	var retrievedUser models.SysUser
	db.First(&retrievedUser, "user_name = ?", "admin")

	// 验证密码
	if !models.VerifyPassword(retrievedUser.PassWord, "admin123") {
		t.Fatalf("password does not match")
	}

}
