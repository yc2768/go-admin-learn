package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50);not null" json:"username"`
	PassWord string `gorm:"type:varchar(255);not null" json:"password"`
	Phone    string `gorm:"type:varchar(20);not null" json:"phone"`
	Avatar   string `gorm:"type:varchar(255);" json:"avatar"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

func CreateSysUser(db *gorm.DB, user *SysUser) error {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PassWord = string(hashedPassWord)
	return db.Create(user).Error
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GetSysUser(name string, password string) (*SysUser, error) {
	data := new(SysUser)
	err := DB.Where("user_name = ?", name).First(data).Error
	if err != nil {
		return nil, err
	}

	// 验证密码
	if !VerifyPassword(data.PassWord, password) {
		return nil, fmt.Errorf("password does not match")
	}

	return data, err
}
