package command

import (
	"fmt"
	"io/ioutil"
	data "web/models/redis"

	yaml "gopkg.in/yaml.v2"

	// "github.com/garyburd/redigo/redis"
	"github.com/NewTrident/baidu"
)

type BdConfig struct {
	AdWidth          int32                                         `yaml:"adslot_width"`
	AdHeight         int32                                         `yaml:"adslot_height"`
	CreativeType     []int32                                       `yaml:"creative_type"`
	CreativeDescType []baidurtb.BidRequest_AdSlot_CreativeDescType `yaml:"creative_desctype"`
	MinimuCpm        int32                                         `yaml:"minimumcpm"`
	AdslotLevel      baidurtb.BidRequest_AdSlot_AdSlotLevel        `yaml:"adslot_level"`
	Platform         baidurtb.BidRequest_Mobile_OS                 `yaml:"platform"`
	Province         string                                        `yaml:"province"`
	City             string                                        `yaml:"city"`
}

func SetBaidu(add string) {

	buf, err := ioutil.ReadFile(add+"baidu.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config BdConfig
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.AdWidth != 0 {
		w := int64(config.AdWidth)
		data.HsetNum("baidu", "width", w)
	}
	if config.AdHeight != 0 {
		h := int64(config.AdHeight)
		data.HsetNum("baidu", "height", h)
	}
	if config.CreativeType != nil {
		data.HsetArr("baidu", "creativetype", config.CreativeType)
	}
	if config.CreativeDescType != nil {
        var desc []int32
		for i:=range config.CreativeDescType{
			if(config.CreativeDescType[i]==baidurtb.BidRequest_AdSlot_STATIC_CREATIVE){
				desc=append(desc,1)
			}else if(config.CreativeDescType[i]==baidurtb.BidRequest_AdSlot_DYNAMIC_CREATIVE){
				desc=append(desc,2)
			}
		}
		// fmt.Println(desc)
		data.HsetArr("baidu", "desctype", desc)
	}
	if config.MinimuCpm != 0 {
		mini := int64(config.MinimuCpm)
		data.HsetNum("baidu", "minimucpm", mini)
	}
	if config.AdslotLevel!=0{
		var level int64
		if(config.AdslotLevel==baidurtb.BidRequest_AdSlot_UNKNOWN_ADB_LEVEL){
			level=0
		}
		if(config.AdslotLevel==baidurtb.BidRequest_AdSlot_TOP){
			level=1
		}
		if(config.AdslotLevel==baidurtb.BidRequest_AdSlot_MED){
			level=2
		}
		if(config.AdslotLevel==baidurtb.BidRequest_AdSlot_TAIL){
			level=3
		}
		if(config.AdslotLevel==baidurtb.BidRequest_AdSlot_LOW){
			level=4
		}
		data.HsetNum("baidu","level",level)
	}
	if(config.Platform!=0){
		var os int64
		if(config.Platform==baidurtb.BidRequest_Mobile_UNKNOWN_OS){
			os=0
		}
		if(config.Platform==baidurtb.BidRequest_Mobile_IOS	){
			os=1
		}
		if(config.Platform==baidurtb.BidRequest_Mobile_ANDROID){
			os=2
		}
		if(config.Platform==baidurtb.BidRequest_Mobile_WINDOWS_PHONE){
			os=3
		}
		data.HsetNum("baidu","os",os)
	}
	if(config.Province!=""){
		data.HsetStr("baidu","province",config.Province)
	}
	if(config.City!=""){
		data.HsetStr("baidu","city",config.City)
	}

}
