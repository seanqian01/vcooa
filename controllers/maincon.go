package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

type ComController struct {
	beego.Controller
}

func (c *ComController) Get() {
	c.TplName = "Merchant.html"
}
