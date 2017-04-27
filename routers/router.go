package routers

import (
	"github.com/goopher503/web/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/login",&controllers.LoginController{},"get,post:Login")

	beego.Router("/admin",&controllers.ADminController{},"post:ADmin")
	beego.Router("/admin",&controllers.AdminController{},"get:Admin")

    beego.Router("/gorun",&controllers.GoRunController{},"post:GoRun")
	beego.Router("/gorun",&controllers.GorunController{},"get:Gorun")

	// beego.Router("/youku",&controllers.YkController{},"get:Youku")
	// beego.Router("/youku",&controllers.YKController{},"post:YouKu")

	// beego.Router("/baidu",&controllers.BaiduController{},"get:Baidu")
	// beego.Router("/baidu",&controllers.BaiDuController{},"post:BaiDu")

	// beego.Router("/smaato",&controllers.SmaatoController{},"get:Smaato")
	// beego.Router("/smaato",&controllers.SmaaToController{},"post:SmaaTo")

	// beego.Router("/mobfox",&controllers.MobfoxController{},"get:Mobfox")
	// beego.Router("/mobfox",&controllers.MobFoxController{},"post:MobFox")
}