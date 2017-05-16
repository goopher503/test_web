package printlog

import (
	"fmt"
    "io/ioutil"
	"os"
	"time"

	// "github.com/NewTrident/openrtb"
)

func SmaatoPrintLog(request []byte,reponse []byte ) {
		err := os.MkdirAll("models/smaato/log/log", 0777)
		if err != nil {
			fmt.Printf("%s\r\n", err.Error())
			os.Exit(-1)
		}
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)

	var filename1 string
	var filename2 string

	filename1 = "models/smaato/log/log/" + tm.Format("2006_01_02_15_04_05") +"_req"+".json"
	filename2 = "models/smaato/log/log/" + tm.Format("2006_01_02_15_04_05") +"_res"+".json"

	logfile1, err := os.OpenFile(filename1, os.O_CREATE|os.O_APPEND, 0664) //创建文件

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile1.Close()

	err = ioutil.WriteFile(filename1, request, 0777)

    if err != nil {

        fmt.Println(err)

    }
	
	logfile2, err := os.OpenFile(filename2, os.O_CREATE|os.O_APPEND, 0664) //追加response

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
    defer logfile2.Close()

	err = ioutil.WriteFile(filename2, reponse, 0777)

    if err != nil {

        fmt.Println(err)

    }

}

