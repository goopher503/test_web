package models

import (
    "github.com/astaxie/beego"

)


func SiteUrl(baseName string) string {
	return beego.AppConfig.String("siteUrl") + "/" + baseName
}
