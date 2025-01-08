package main

import (
	"go-admin/models"
	"go-admin/router"
)

func main() {

	models.NewGormDB()

	r := router.App()
	r.Run(":8080")
}
