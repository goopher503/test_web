package command

import (
	"fmt"
	"io/ioutil"
	data "web/models/redis"

	yaml "gopkg.in/yaml.v2"

	// "github.com/garyburd/redigo/redis"
)

type YkConfig struct {
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

func SetYouku(add string) {

	buf, err := ioutil.ReadFile(add+"youku.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config YkConfig
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.Width != 0 {
		w := int64(config.Width)
		data.HsetNum("youku", "width", w)
	}
	if config.Height != 0 {
		h := int64(config.Height)
		data.HsetNum("youku", "height", h)
	}

	if(config.Os!=""){
		data.HsetStr("youku","os",config.Os)
	}
	if(config.Osv!=""){
		data.HsetStr("youku","osv",config.Osv)
	}

}
