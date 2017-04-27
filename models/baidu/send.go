package baidu

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/goopher503/web/models/baidu/log"
	"github.com/goopher503/web/models/baidu/request"
	"github.com/goopher503/web/models/descryptbes"

	rlib "github.com/garyburd/redigo/redis"

	// "strconv"
	"strings"

	"github.com/NewTrident/baidu"
	httpclient "github.com/go-httpclient"
	"github.com/golang/protobuf/proto"
)

var address string = "127.0.0.1:6379"
var pass string = "molibaiju"

func Send(num int64, ip string) *baidurtb.BidResponse {

	var r request.Request
	var response baidurtb.BidResponse
	var i int64

	transport := &httpclient.Transport{
		ConnectTimeout:        10 * time.Second,
		RequestTimeout:        30 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
	}
	defer transport.Close()

	client := &http.Client{Transport: transport}

	for i = 0; i < num; i++ {
		cli, _ := rlib.Dial("tcp", address, rlib.DialPassword(pass))
		defer cli.Close()
		flag, _ := rlib.Int(cli.Do("GET", "numflag"))
		// fmt.Println("flag:",flag)
		if flag == 1 {
			break
		}
		r.MakeBesRequest()
		requestbyte, _ := proto.Marshal(r.BidRequest)
		// requestjson,_:=json.Marshal(r.BidRequest)
		body := bytes.NewReader(requestbyte)
		var url, url2 string
		if ip != "" {
			url = "http://" + ip + "/bes"
			url2 = "http://" + ip + "/bespv"
		} else {
			url = "http://localhost:8000/bes"
			url2 = "http://localhost:8000/bespv"
		}
		req, _ := http.NewRequest("POST", url, body)
		t1 := time.Now()
		resp, err := client.Do(req)
		elapsed := time.Since(t1)
		fmt.Println("App elapsed: ", elapsed)
		// respcode := resp.StatusCode
		// fmt.Println("httpcode:",respcode)
		if err != nil {
			fmt.Println("err:", err)
		}

		responsebody, err := ioutil.ReadAll(resp.Body)
		// responsebodyjson, err := ioutil.ReadAll(resp.Body)
		var str string
		var star, end int
		proto.Unmarshal(responsebody, &response)
		for i := range response.GetAd() {
			if response.GetAd()[i] != nil {
				res2 := response.GetAd()[i].GetHtmlSnippet()
				for i := range res2 {
					if res2[i] == 98 && res2[i+1] == 101 && res2[i+2] == 115 && res2[i+3] == 112 && res2[i+4] == 118 {
						star = i + 6
						for j := i + 5; ; j++ {
							if res2[j] == 34 {
								end = j
								break
							}
						}
						break
					}
				}
				str = res2[star:end]
				str = url2 + "?" + str
				price := *response.GetAd()[i].MaxCpm
				pri := descryptbes.Encrypt(price)

				str = strings.Replace(str, "%%PRICE%%", pri, -1)
				// fmt.Println("str:",str)
				respp, _ := http.Get(str)
				fmt.Println("bes_win_code", i+1, ":", respp.StatusCode)
			}
		}
		// json.Unmarshal(responsebodyjson,&response)
		i := 1
		printlog.BaiDuPrintLog(r.BidRequest, &response, i)
		defer resp.Body.Close()
	}
	return &response
}
