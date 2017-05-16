package smaato

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
	"strconv"
	"strings"
	"time"
	"github.com/goopher503/web/models/smaato/log"
	"github.com/goopher503/web/models/smaato/request"
	rlib "github.com/garyburd/redigo/redis"

	"github.com/NewTrident/openrtb"
	httpclient "github.com/mreiferson/go-httpclient"
)
var address string = "127.0.0.1:6379"
var pass string = "molibaiju"
func Send(num int64,ip string) *openrtb.BidResponse {

	var r request.Request
	var response openrtb.BidResponse
	var i int64

	transport := &httpclient.Transport{
		ConnectTimeout:        10 * time.Second,
		RequestTimeout:        10 * time.Second,
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
		r.MakeRequest()
		requestbyte, _ := r.BidRequest.MarshalJSON()
		body := bytes.NewReader(requestbyte)
        var url,url2,str string
		if(ip!=""){
			url = "http://"+ip+"/smaato"
			url2 = ip
		}else{
			url = "http://localhost:8000/smaato"
			url2 = "localhost:8000"
		}
		req, _ := http.NewRequest("POST", url, body)
		resp, err := client.Do(req)
		respcode := resp.StatusCode
		// fmt.Println("smaato_httpcode:",respcode)

		if err != nil {
			fmt.Println("err:", err)
		}
		
		responsebody, err := ioutil.ReadAll(resp.Body)
		err = response.UnmarshalJSON(responsebody)
		printlog.SmaatoPrintLog(requestbyte, responsebody)
		defer resp.Body.Close()

		if(respcode==200){
            for i:= range response.SeatBid{
				str = response.SeatBid[i].Bid[i].NURL
				price := response.SeatBid[i].Bid[i].Price
				pri := strconv.FormatFloat(price,'f',-1,64)
				str = strings.Replace(str,"uscount.smartkeeplive.com",url2,-1)
				str = strings.Replace(str,"${AUCTION_PRICE}",pri,-1)
                // fmt.Println("smaato_str",str)
				respp,_ := http.Get(str)
				fmt.Println("smaato_win_code",i+1,":",respp.StatusCode)
			}
		}
	}
	return &response
}
