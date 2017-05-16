package youku

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
	"strconv"
	"strings"
	"time"
	base64old "encoding/base64"
	"github.com/goopher503/web/models/youku/log"
	"github.com/goopher503/web/models/youku/request"
	rlib "github.com/garyburd/redigo/redis"
	"github.com/goopher503/web/models/decryptyk"

	"github.com/NewTrident/youku"
	httpclient "github.com/mreiferson/go-httpclient"
)
var address string = "127.0.0.1:6379"
var pass string = "molibaiju"
func Send(num int64,ip string) *youku.BidResponse {

	var r request.Request
	var response youku.BidResponse
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
		r.MakeYkRequest()
		requestbyte, _ := r.BidRequest.MarshalJSON()
		body := bytes.NewReader(requestbyte)
        var url,url2,str string
		if(ip!=""){
			url = "http://"+ip+"/youku"
			url2 = ip
		}else{
			url = "http://localhost:8000/youku"
			url2 = "localhost:8000"
		}
		req, _ := http.NewRequest("POST", url, body)
		resp, err := client.Do(req)
        respcode := resp.StatusCode
		// fmt.Println("youku_httpcode:",respcode)

		if err != nil {
			fmt.Println("err:", err)
		}

		responsebody, err := ioutil.ReadAll(resp.Body)
		err = response.UnmarshalJSON(responsebody)
		printlog.YkPrintLog(requestbyte, responsebody)
		defer resp.Body.Close()

		if(respcode==200){
            for i:= range response.SeatBid{
				str = response.SeatBid[i].Bid[i].NUrl
				price := response.SeatBid[i].Bid[i].Price
				pri := strconv.FormatFloat(price,'f',-1,64)

				keyStr := "c2de1527874d4780a1544cfc80c0a27b"
				pri = AesEncPrice(pri,keyStr)


				str = strings.Replace(str,"bid.cpxengine.com",url2,-1)
				str = strings.Replace(str,"${AUCTION_PRICE}",pri,-1)
                // fmt.Println("youku_str",str)
				respp,_ := http.Get(str)
				fmt.Println("youku_win_code",i+1,":",respp.StatusCode)
			}
		}
	}
	return &response

}

func MakeKey(input []byte) []byte {

	byte2 := make([]byte, len(input))
	for i, j := range input {
		if j < 58 {
			byte2[i] = j - 48
		} else {
			byte2[i] = j - 87
		}
	}

	output := make([]byte, len(byte2)/2)
	for i := 0; i < len(byte2); i = i + 2 {
		output[i/2] = byte2[i]<<4 + byte2[i+1]
	}
	return output
}

func AesEncPrice(price, keystr string) string {

	key := MakeKey([]byte(keystr))
    
	pri := []byte(price)
	outputbyte, err := decryptyk.AesEncrypt(pri,key)
	if err != nil {
		return ""	
	}

	output := base64old.RawURLEncoding.EncodeToString(outputbyte)

	return output
}
