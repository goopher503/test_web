package command

import (
	"fmt"
	"io/ioutil"
	data "web/models/redis"

	yaml "gopkg.in/yaml.v2"

	// "github.com/garyburd/redigo/redis"
)

type MfConfig struct {
	Width     uint64  `yaml:"width"`
	Height    uint64  `yaml:"height"`
	Os        string  `yaml:"os"`
	Osv       string  `yaml:"osv"`
	Country   string  `yaml:"country"`
	Region    string  `yaml:"region"`
	City      string  `yaml:"city"`
	At        int8    `yaml:"at"`
	BidFloor  float64 `yaml:"bidfloor"`
	PublishID string  `yaml:"publishid"`
}

func SetMobfox(add string) {

	buf, err := ioutil.ReadFile(add+"mobfox.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config MfConfig
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.Width != 0 {
		w := int64(config.Width)
		data.HsetNum("mobfox", "width", w)
	}
	if config.Height != 0 {
		h := int64(config.Height)
		data.HsetNum("mobfox", "height", h)
	}

	if(config.Country!=""){
		data.HsetStr("mobfox","os",config.Os)
	}
		if(config.Osv!=""){
		data.HsetStr("mobfox","osv",config.Osv)
	}
	if(config.Country!=""){
		data.HsetStr("mobfox","country",config.Country)
	}
	if(config.Region!=""){
		data.HsetStr("mobfox","region",config.Region)
	}
	if(config.City!=""){
		data.HsetStr("mobfox","city",config.City)
	}

	if config.At != 0 {
		at := int64(config.At)
		data.HsetNum("mobfox", "at", at)
	}
	if config.BidFloor != 0 {
		data.HsetFlo("mobfox", "bidfloor", config.BidFloor)
	}
	if(config.PublishID!=""){
		data.HsetStr("mobfox","publishid",config.PublishID)
	}


}
