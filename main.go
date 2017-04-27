package main

import (
	_ "web/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {

	_, ok := ctx.Input.Session("userlogin").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}


