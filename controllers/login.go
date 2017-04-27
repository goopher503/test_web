package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	
	"fmt"
)

type LoginController struct {
	beego.Controller
}

	
func (this *LoginController) Login() {
	
	user:=this.GetString("user")
	pass:=this.GetString("password")
	// action := this.GetString("action")
    this.DelSession("userlogin")
	if(user=="admin"&&pass=="admin"){
		this.SetSession("userlogin",fmt.Sprintf(user))
		// fmt.Println("session:",this.GetSession("userlogin"))
		this.Redirect(models.SiteUrl("admin"),302)
	}else{
		 this.TplName="login.html"
	}
}




