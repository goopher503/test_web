package platform

import (

    "github.com/goopher503/web/models/youku"
    data "github.com/goopher503/web/models/redis"
	you "github.com/NewTrident/youku"

	// "fmt"
)

func YouKu(width int64,height int64,os string,osv string,num int64,ip string) string {

    if width != 0 {
		data.HsetNum("youku", "width", width)
	}
	if height != 0 {
		data.HsetNum("youku", "height", height)
	}
	if(os!=""){
		data.HsetStr("youku","os",os)
	}
    if(osv!=""){
		data.HsetStr("youku","osv",osv)
	}

	var resp *you.BidResponse
	var adm string
    resp=youku.Send(num,ip)
	// fmt.Println("response:",resp)
    if(num==1){
		adm= resp.SeatBid[0].Bid[0].Adm
	}else{
		adm=""
	}
	return adm
}

