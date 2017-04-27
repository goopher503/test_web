package command

import (
	"fmt"
	"io/ioutil"
	data "github.com/goopher503/web/models/redis"

	yaml "gopkg.in/yaml.v2"

	// "github.com/garyburd/redigo/redis"
)

type StConfig struct {
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

func SetSmaato(add string) {

	buf, err := ioutil.ReadFile(add+"smaato.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config StConfig
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.Width != 0 {
		w := int64(config.Width)
		data.HsetNum("smaato", "width", w)
	}
	if config.Height != 0 {
		h := int64(config.Height)
		data.HsetNum("smaato", "height", h)
	}

	if(config.Os!=""){
		data.HsetStr("smaato","os",config.Os)
	}
	if(config.Osv!=""){
		data.HsetStr("smaato","osv",config.Osv)
	}
	if(config.Country!=""){
		data.HsetStr("smaato","country",config.Country)
	}
	if(config.Region!=""){
		data.HsetStr("smaato","region",config.Region)
	}
	if(config.City!=""){
		data.HsetStr("smaato","city",config.City)
	}

	if config.At != 0 {
		at := int64(config.At)
		data.HsetNum("smaato", "at", at)
	}
	if config.BidFloor != 0 {
		data.HsetFlo("smaato", "bidfloor", config.BidFloor)
	}
	if(config.PublishID!=""){
		data.HsetStr("smaato","publishid",config.PublishID)
	}


}
