package platform

import (

    data "github.com/goopher503/web/models/redis"
    "github.com/goopher503/web/models/smaato"
	"github.com/NewTrident/openrtb"

	"fmt"
)

func SmaaTo(width int64, height int64, os string,osv string,country string,region string, city string, bidfloor float64,id string,num int64,ip string) string {

    var err error
    if width != 0 {
		data.HsetNum("smaato", "width", width)
	}
	if height != 0 {
		data.HsetNum("smaato", "height", height)
	}
	if bidfloor != 0 {
		err=data.HsetFlo("smaato", "bidfloor", bidfloor)
	}
	if(os!=""){
		data.HsetStr("smaato","os",os)
	}
    if(osv!=""){
		data.HsetStr("smaato","osv",osv)
	}
    if(country!=""){
		data.HsetStr("smaato","country",country)
	}
	if(region!=""){
		data.HsetStr("smaato","region",region)
	}
	if(city!=""){
		data.HsetStr("smaato","city",city)
	}
    if(id!=""){
        data.HsetStr("smaato","publishid",id)
    }

	if(err!=nil){
		fmt.Println("set error:",err)
	}
     
    var resp *openrtb.BidResponse
	var adm string
	resp = smaato.Send(num,ip)
	if num == 1 {
		adm = resp.SeatBid[0].Bid[0].AdM
	}else{
		adm = ""
	}
	return adm
}
