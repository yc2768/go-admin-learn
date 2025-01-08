package service

import (
	"github.com/gin-gonic/gin"
	"go-admin/define"
	"go-admin/helper"
	"go-admin/models"
	"net/http"
)

func Login(c *gin.Context) {
	in := new(LoginPassWordRequest)
	err := c.ShouldBindJSON(&in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "解析Json错误!",
		})
		return
	}

	sysUser, err := models.GetSysUser(in.UserName, in.PassWord)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err := helper.GenerateToken(sysUser.ID, sysUser.UserName, define.TokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}

	refreshToken, err := helper.GenerateToken(sysUser.ID, sysUser.UserName, define.RefreshTokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}

	data := &LoginPassWordResponse{
		Toekn:        token,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": data,
	})

}
