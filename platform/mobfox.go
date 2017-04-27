package platform

import (
	"github.com/goopher503/web/models/mobfox"
	data "github.com/goopher503/web/models/redis"

	"github.com/NewTrident/mobfox"
)

func MobFox(width int64, height int64, os string, osv string, country string, region string, city string, bidfloor float64, id string,num int64,ip string) string {

	if width != 0 {
		data.HsetNum("mobfox", "width", width)
	}
	if height != 0 {
		data.HsetNum("mobfox", "height", height)
	}
	if bidfloor != 0 {
		data.HsetFlo("mobfox", "bidfloor", bidfloor)
	}
	if os != "" {
		data.HsetStr("mobfox", "os", os)
	}
	if osv != "" {
		data.HsetStr("mobfox", "osv", osv)
	}
	if country != "" {
		data.HsetStr("mobfox", "country", country)
	}
	if region != "" {
		data.HsetStr("mobfox", "region", region)
	}
	if city != "" {
		data.HsetStr("mobfox", "city", city)
	}
	if id != "" {
		data.HsetStr("mobfox", "publishid", id)
	}

	var resp *openrtb.BidResponse
	var adm string
	resp = mobfox.Send(num,ip)
	if num == 1 {
		adm = resp.SeatBid[0].Bid[0].AdM
	}else{
		adm = ""
	}
	return adm
}
