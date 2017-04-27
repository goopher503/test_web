package platform

import (
	"github.com/goopher503/web/models/baidu"
	data "github.com/goopher503/web/models/redis"
	"github.com/NewTrident/baidu"
	// "encoding/json"
)

func Baidu(width int64, height int64, os string, region string, city string, bidfloor float64,num int64,ip string) string{

	var oos int64
	switch os {
	case "iOS":
		oos = 1
	case "Android":
		oos = 2
	case "WindowsPhone":
		oos = 3
	default:
		oos = 0
	}
	if width != 0 {
		data.HsetNum("baidu", "width", width)
	}
	if height != 0 {
		data.HsetNum("baidu", "height", height)
	}
	if oos != 0 {
		data.HsetNum("baidu", "os", oos)
	}
	if region != "" {
		data.HsetStr("baidu", "province", region)
	}
	if city != "" {
		data.HsetStr("baidu", "city", city)
	}
	if bidfloor != 0 {
		minimucpm := int64(bidfloor)
		data.HsetNum("baidu", "minimucpm", minimucpm)
	}
    
    var resp *baidurtb.BidResponse
	var adm string
	resp=baidu.Send(num,ip)
	if(num==1){
		adm=resp.Ad[0].GetHtmlSnippet()
	}else{
		adm=""
	}
	return adm	
}
