package redis

import (
	// "fmt"
	// "time"

	// log "github.com/cihub/seelog"
	rlib "github.com/garyburd/redigo/redis"

	// "github.com/NewTrident/GoDSP/common/consts"
)

var add string = "127.0.0.1:6379"
var pass string = "molibaiju"



func HsetNum(platfrom string,key string,num int64) error {
	
	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))

	defer cli.Close()

	_,err = cli.Do("HSET",platfrom,key,num)
	return err
}

func HgetNum(platfrom string,key string) (int64,error) {
	
	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))
	defer cli.Close()

	num,err := rlib.Int64(cli.Do("HGET",platfrom,key))
	return num, err
}

func HsetStr(platfrom string,key string,str string) error {
	
	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))
	defer cli.Close()

	_,err = cli.Do("HSET",platfrom,key,str)
	return err
}

func HgetStr(platfrom string,key string) (string,error) {

	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))
	defer cli.Close()

	str,err := rlib.String(cli.Do("HGET",platfrom,key))
	return str,err
}

func HsetFlo(platfrom string,key string,num float64) error {
	
	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))
	defer cli.Close()

	_,err = cli.Do("HSET",platfrom,key,num)
	return err
}

func HgetFlo(platfrom string,key string) (float64,error) {
	
	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))
	defer cli.Close()

	num,err := rlib.Float64(cli.Do("HGET",platfrom,key))
	return num, err
}

func HsetArr(platfrom string,key string,arr []int32) error {
	
	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))
	defer cli.Close()
    for i:=range arr{
	   _,err := cli.Do("HSET",platfrom,key+string(i+49),arr[i])
	   if(err!=nil){
		   return err
	   }
	}
	return err
}

func HgetArr(platfrom string,key string) ([]int32,error) {
	
	var err error
	cli, err := rlib.Dial("tcp", add, rlib.DialPassword(pass))
	defer cli.Close()
	var arrr [8]int64
	var arr []int32
	err = nil
    for i:=0;;i++{
	  arrr[i],err = rlib.Int64(cli.Do("HGET",platfrom,key+string(i+49)))
	  if(arrr[i]!=0){
         arr=append(arr,int32(arrr[i]))
	  }else{
		  break
	  }
	}
	return arr,err
}

