package request

import (
	// "fmt"
	// "io/ioutil"

	"github.com/pborman/uuid"

	rlib "web/models/redis"

	// yaml "gopkg.in/yaml.v2"
)

// type Config struct {
// 	// BidFloor float64 `yaml:"bidfloor"`
// 	Width     uint64  `yaml:"width"`
// 	Height    uint64  `yaml:"height"`
// 	Os        string  `yaml:"os"`
// 	Osv       string  `yaml:"osv"`
// 	Country   string  `yaml:"country"`
// 	Region    string  `yaml:"region"`
// 	City      string  `yaml:"city"`
// 	At        int8    `yaml:"at"`
// 	BidFloor  float64 `yaml:"bidfloor"`
// 	PublishID string  `yaml:"publishid"`
// }

func (r *Request) FillMobRequest() error {

	// buf, err := ioutil.ReadFile("models//mobfox//request//data.yml")
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

    var width,height,at int64
	var country,region,city,os,osv,publishid string
	var bidfloor float64
	var err error

	//从redis中读取数据
	width,err=rlib.HgetNum("mobfox","width")
	height,err=rlib.HgetNum("mobfox","height")
	at,err=rlib.HgetNum("mobfox","at")
	country,err=rlib.HgetStr("mobfox","country")
	region,err=rlib.HgetStr("mobfox","region")
	city,err=rlib.HgetStr("mobfox","city")
	os,err=rlib.HgetStr("mobfox","os")
	osv,err=rlib.HgetStr("mobfox","osv")
	publishid,err=rlib.HgetStr("mobfox","publishid")
	bidfloor,err=rlib.HgetFlo("mobfox","bidfloor")


	//填充拍卖类型
	if at != 0 {
		a:=int8(at)
		r.BidRequest.AT = a
	}

	//填充竞拍最低价
	if bidfloor != 0 {
		r.BidRequest.Imp[0].BidFloor = bidfloor
	}

	//填充imp的尺度
	if width != 0 && height != 0 {
		for i := range r.BidRequest.Imp {
			w:=uint64(width)
			h:=uint64(height)
			r.BidRequest.Imp[i].Banner.W = w
			r.BidRequest.Imp[i].Banner.H = h
		}
	}

	//填充AppID，用于黑白名单的判断
	if r.BidRequest.App != nil && publishid != "" {
		r.BidRequest.App.ID = publishid
	}
	if r.BidRequest.Site != nil && publishid != "" {
		r.BidRequest.Site.ID = publishid
	}

	//填充操作系统
	if os != "" {
		r.BidRequest.Device.OS = os
		if os == "iOS" {
			r.BidRequest.Ext.Udi.IDFA = uuid.NewRandom().String()
			r.BidRequest.Ext.Udi.GAID = ""
		}
		if os == "Android" {
			r.BidRequest.Ext.Udi.IDFA = ""
			r.BidRequest.Ext.Udi.GAID = uuid.NewRandom().String()
		}

	}

	//填充操作系统版本
	if osv != "" {
		r.BidRequest.Device.OSV = osv
	}

	//填充geo
	if country != "" && region != "" && city != "" {
		r.BidRequest.Device.Geo.Country = country
		r.BidRequest.Device.Geo.Region = region
		r.BidRequest.Device.Geo.City = city
	}
	return err
}
