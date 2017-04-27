package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	rlib "github.com/garyburd/redigo/redis"
)

type GorunController struct {
	beego.Controller
}

type GoRunController struct {
	beego.Controller
}

func (this *GorunController) Gorun() {

	this.TplName = "gorun.html"

}

func (this *GoRunController) GoRun() {

	var address string = "127.0.0.1:6379"
	var pass string = "molibaiju"
	cli, _ := rlib.Dial("tcp", address, rlib.DialPassword(pass))
	defer cli.Close()
	_, _ = cli.Do("SET", "numflag", 1)

	this.Redirect(models.SiteUrl("admin"), 302)

}
