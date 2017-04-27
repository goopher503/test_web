package request

import (
	// "fmt"
	// "io/ioutil"

	// yaml "gopkg.in/yaml.v2"
	rlib "github.com/goopher503/web/models/redis"

	"github.com/NewTrident/baidu"
)

// type Config struct {
// 	AdWidth          int32                                         `yaml:"adslot_width"`
// 	AdHeight         int32                                         `yaml:"adslot_height"`
// 	CreativeType     []int32                                       `yaml:"creative_type"`
// 	CreativeDescType []baidurtb.BidRequest_AdSlot_CreativeDescType `yaml:"creative_desctype"`
// 	MinimuCpm        int32                                         `yaml:"minimumcpm"`
// 	AdslotLevel      baidurtb.BidRequest_AdSlot_AdSlotLevel        `yaml:"adslot_level"`
// 	Platform         baidurtb.BidRequest_Mobile_OS                 `yaml:"platform"`
// 	Province         string                                        `yaml:"province"`
// 	City             string                                        `yaml:"city"`
// }

func (r *Request) FillBesRequest() error {

	// buf, err := ioutil.ReadFile("models//baidu//request//data.yml")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// var config Config
	// err = yaml.Unmarshal(buf, &config)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	var width,height,minimucpm,level,os int64
	var province,city string
	var creativetype,desctype []int32
	var err error

	//从redis中读取数据
	width,err=rlib.HgetNum("baidu","width")
	height,err=rlib.HgetNum("baidu","height")
	creativetype,err=rlib.HgetArr("baidu","creativetype")
	desctype,err=rlib.HgetArr("baidu","desctype")
	minimucpm,err=rlib.HgetNum("baidu","minimucpm")
	level,err=rlib.HgetNum("baidu","level")
	os,err=rlib.HgetNum("baidu","os")
	province,err=rlib.HgetStr("baidu","province")
	city,err=rlib.HgetStr("baidu","city")


	//填充adslot
	for i := range r.BidRequest.Adslot {
		if width != 0 && height != 0 {
            w:=int32(width)
			h:=int32(height)
			r.BidRequest.Adslot[i].Width = &w
			r.BidRequest.Adslot[i].Height = &h
		}
		for j:=range creativetype{
			if(creativetype[j]!=0){
				r.BidRequest.Adslot[i].CreativeType[j]=creativetype[j]
			}
		}
		for j:=range desctype{
			if(desctype[j]!=0){
				switch desctype[j] {
				case 1:
					r.BidRequest.Adslot[i].CreativeDescType=append(r.BidRequest.Adslot[i].CreativeDescType,baidurtb.BidRequest_AdSlot_STATIC_CREATIVE)
				case 2:
				    r.BidRequest.Adslot[i].CreativeDescType=append(r.BidRequest.Adslot[i].CreativeDescType,baidurtb.BidRequest_AdSlot_DYNAMIC_CREATIVE)
				}				
			}
		}
		if minimucpm != 0 {
			mini:=int32(minimucpm)
			r.BidRequest.Adslot[i].MinimumCpm = &mini
		}
		if level != 0 {
			var le baidurtb.BidRequest_AdSlot_AdSlotLevel
			switch level {
			case 0:
			    le=baidurtb.BidRequest_AdSlot_UNKNOWN_ADB_LEVEL
			case 1:
			    le=baidurtb.BidRequest_AdSlot_TOP
			case 2:
			    le=baidurtb.BidRequest_AdSlot_MED
			case 3:
			    le=baidurtb.BidRequest_AdSlot_TAIL
			case 4:
			    le=baidurtb.BidRequest_AdSlot_LOW	
			}
            r.BidRequest.Adslot[i].AdslotLevel=&le
		}
	}

	//填充mobile
	if os != 0 {
		var oos baidurtb.BidRequest_Mobile_OS
		switch os {
		case 0:
			oos=baidurtb.BidRequest_Mobile_UNKNOWN_OS
		case 1:
			oos=baidurtb.BidRequest_Mobile_IOS
		case 2:
			oos=baidurtb.BidRequest_Mobile_ANDROID
		case 3:
			oos=baidurtb.BidRequest_Mobile_WINDOWS_PHONE
		}
		r.BidRequest.Mobile.Platform = &oos
	}
	if r.BidRequest.GetMobile().GetPlatform() == baidurtb.BidRequest_Mobile_IOS {
		for i := range r.BidRequest.Mobile.ForAdvertisingId {
			*r.BidRequest.Mobile.ForAdvertisingId[i].Type = baidurtb.BidRequest_Mobile_ForAdvertisingID_IDFA
		}
	}
	if r.BidRequest.GetMobile().GetPlatform() == baidurtb.BidRequest_Mobile_ANDROID {
		for i := range r.BidRequest.Mobile.ForAdvertisingId {
			*r.BidRequest.Mobile.ForAdvertisingId[i].Type = baidurtb.BidRequest_Mobile_ForAdvertisingID_ANDROID_ID
		}
	}
	if province != "" {
		r.BidRequest.UserGeoInfo.UserLocation.Province = &province
	}
	if city != "" {
		r.BidRequest.UserGeoInfo.UserLocation.City = &city
	}
	return err
}
