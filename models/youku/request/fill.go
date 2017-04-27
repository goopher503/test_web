package request

import (
	// "fmt"
	// "io/ioutil"
	"github.com/pborman/uuid"
	rlib "github.com/goopher503/web/models/redis"

	// yaml "gopkg.in/yaml.v2"
)

type Config struct {
	// BidFloor float64 `yaml:"bidfloor"`
	Width  uint64 `yaml:"width"`
	Height uint64 `yaml:"height"`
	Os     string `yaml:"os"`
	Osv    string `yaml:"osv"`
	// Country  string  `yaml:"country"`
	// Region   string  `yaml:"region"`
	// City     string  `yaml:"city"`
	// At       int8    `yaml:"at"`
	// BidFloor float64 `yaml:"bidfloor"`
}

func (r *Request) FillYkRequest() error {

	// buf, err := ioutil.ReadFile("models//youku//request//data.yml")
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
    var width,height int64 
	var os,osv string
	var err error

	//从redis中读取数据
	width,err=rlib.HgetNum("youku","width")
	height,err=rlib.HgetNum("youku","height")
	os,err=rlib.HgetStr("youku","os")
	osv,err=rlib.HgetStr("youku","osv")

	//填充尺度
	if width != 0 && height != 0 {
		for i := range r.BidRequest.Imp {
			w:=uint64(width)
			h:=uint64(height)
			r.BidRequest.Imp[i].Banner.W = w
			r.BidRequest.Imp[i].Banner.H = h
		}
	}

	//填充操作系统
	if os != "" {
		r.BidRequest.Device.OS = os
		if os == "iOS" {
			r.BidRequest.Device.IDFA= uuid.NewRandom().String()
			r.BidRequest.Device.AndroidID = ""
		}
		if os == "Android" {
			r.BidRequest.Device.IDFA= ""
			r.BidRequest.Device.AndroidID =uuid.NewRandom().String()
		}

	}
   
    //填充操作系统版本
	if osv != "" {
		r.BidRequest.Device.OSV = osv
	}

	return err
}
