package controllers

import (
	"web/command/create"
	"web/models"
	"web/platform"
	// "runtime"

	rlib "github.com/garyburd/redigo/redis"

	"github.com/astaxie/beego"
	// "web/models/youku"
	"fmt"
)

var address string = "127.0.0.1:6379"
var pass string = "molibaiju"
var c chan int

type AdminController struct {
	beego.Controller
}
type ADminController struct {
	beego.Controller
}

func (this *ADminController) ADmin() {

	cli, _ := rlib.Dial("tcp", address, rlib.DialPassword(pass))
	defer cli.Close()
	_, _ = cli.Do("SET", "numflag", 0)

	// ssp:=this.GetString("SSP")
	ssp := this.Input().Get("ssp")
	num, err := this.GetInt64("num")
	width, err := this.GetInt64("width")
	height, err := this.GetInt64("height")
	os := this.GetString("os")
	osv := this.GetString("osv")
	country := this.GetString("country")
	region := this.GetString("region")
	city := this.GetString("city")
	bidfloor, err := this.GetFloat("bidfloor")
	id := this.GetString("id")
	ip := this.GetString("ip")

	switch ssp {
	case "baidu":
		// fmt.Println(ad)
		if num == 1 || num == 0 {
			num = 1
			adm := platform.Baidu(width, height, os, region, city, bidfloor, num, ip)
			ad := models.Setadm(adm)
			this.Ctx.WriteString(ad)
		} else {
			go func() {
				_ = platform.Baidu(width, height, os, region, city, bidfloor, num, ip)
			}()
			this.Redirect(models.SiteUrl("gorun"), 302)
		}
	case "youku":
		// fmt.Println(ad)
		if num == 1 || num == 0 {
			num = 1
			adm := platform.YouKu(width, height, os, osv, num, ip)
			ad := models.Setadm(adm)
			this.Ctx.WriteString(ad)
		} else {
			go func() {
				_ = platform.YouKu(width, height, os, osv, num, ip)
			}()
			this.Redirect(models.SiteUrl("gorun"), 302)
		}
	case "smaato":
		if num == 1 || num == 0 {
			num = 1
			adm := platform.SmaaTo(width, height, os, osv, country, region, city, bidfloor, id, num, ip)
			ad := models.Setadm(adm)
			this.Ctx.WriteString(ad)
		} else {
			go func() {
				_ = platform.SmaaTo(width, height, os, osv, country, region, city, bidfloor, id, num, ip)
			}()
			this.Redirect(models.SiteUrl("gorun"), 302)
		}
	case "mobfox":
		if num == 1 || num == 0 {
			num = 1
			adm := platform.MobFox(width, height, os, osv, country, region, city, bidfloor, id, num, ip)
			ad := models.Setadm(adm)
			this.Ctx.WriteString(ad)
		} else {
			go func() {
				_ = platform.MobFox(width, height, os, osv, country, region, city, bidfloor, id, num, ip)
			}()
			this.Redirect(models.SiteUrl("gorun"), 302)
		}
	default:
		this.TplName = "admin.html"
	}
	fmt.Println("error:", err)
}
func (this *AdminController) Admin() {

	name := this.GetSession("userlogin")
	// fmt.Println("name:",name)
	if name == nil {
		this.TplName = "login.html"
	} else {
		cli, _ := rlib.Dial("tcp", address, rlib.DialPassword(pass))
		defer cli.Close()
		_, _ = cli.Do("SET", "numflag", 1)

		add := "command//yml//"
		command.SetBaidu(add)
		command.SetMobfox(add)
		command.SetSmaato(add)
		command.SetYouku(add)
		this.TplName = "admin.html"
	}
}
