package controllers

import (
	"github.com/astaxie/beego"
)

type CgangController struct {
	beego.Controller
}
func (p * CgangController) Get() {
	p.TplName = "password.html"
}
