package controllers

import (
	"rtstk/backend/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UsersController struct {
	beego.Controller
}

func (c *UsersController) GetUsers() {
    c.Data["json"] = map[string]any{"users": []models.User{
        {ID: 1, FirstName: "Mika", LastName: "Rael", Email: "mika@artistic.com", Password: "password"},
        {ID: 2, FirstName: "Leen", LastName: "Shaniz", Email: "leen@artistic.com", Password: "password"},
        {ID: 2, FirstName: "Momo", LastName: "Lew", Email: "Lew@artistic.com", Password: "password"},
    }}
    c.ServeJSON()
}
