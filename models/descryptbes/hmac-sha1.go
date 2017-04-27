package descryptbes

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	// "fmt"
)

func Encrypt(priceint int32) string {

	encryption_key := []byte{0x01, 0x5f, 0xdc, 0x74, 0x00, 0xa6, 0xa4, 0x00, 0x04, 0x7c, 0x2f, 0xed, 0x00, 0xa6, 0xa4, 0x00, 0x04, 0x7d, 0x1d, 0x3e, 0x00, 0xa6, 0xa4, 0x00, 0x04, 0x7d, 0x20, 0xaa, 0xdb, 0x01, 0x5c, 0xf8}
	integrity_key_v := []byte{0x01, 0x5f, 0xdc, 0x74, 0x00, 0xa6, 0xa4, 0x00, 0x04, 0x7e, 0x7a, 0x34, 0x00, 0xa6, 0xa4, 0x00, 0x04, 0x7e, 0x7e, 0x96, 0x00, 0xa6, 0xa4, 0x00, 0x04, 0x7e, 0x80, 0x90, 0x9f, 0x07, 0x91, 0x4c}
	iv := []byte{88, 208, 193, 206, 0, 6, 221, 172, 123, 140, 74, 96, 91, 146, 0, 242}
	len_price := 8

	//生成价格字节数组{0,0,0,0,0,0,0,156}
	price := make([]byte, len_price)
	for i := 0; i < 8; i++ {
		j := priceint % 1000
		priceint = priceint / 1000
		price[len_price-i-1] = byte(j)
	}
     // fmt.Println("price:",price)
	
	//将价格与iv数组合成
	conf_sig_iv := make([]byte, len(price)+len(iv))
	for i := 0; i < 8; i++ {
		conf_sig_iv[i] = price[i]
	}
	for i := 0; i < 16; i++ {
		conf_sig_iv[i+8] = iv[i]
	}
    // fmt.Println("conf_sig_iv:",conf_sig_iv)

	//对conf_sig_iv使用integrity_key_v作为key进行hmac-sha1加密，生成sig字节数组
	mac2 := hmac.New(sha1.New, integrity_key_v)
	mac2.Write(conf_sig_iv)
	sig := mac2.Sum(nil)
    // fmt.Println("sig:",sig)
   
    //对iv使用encryption_key作为key进行hmac-sha1加密，生成price_pad字节数组
    mac := hmac.New(sha1.New, encryption_key)
	mac.Write(iv)
	price_pad := mac.Sum(nil)
    // fmt.Println("price_pad:",price_pad)

	//peice_pad与price进行异或，合成p字节数组
    p := make([]byte,len(price))
    for i := range p{
        p[i]=price[i] ^ price_pad[i]
    }
    // fmt.Println("p:",p)

    //将iv,p,sig合成
    encrypt := make([]byte,28)
    for i:=0;i<16;i++{
        encrypt[i]=iv[i]
    }
    for i:=0;i<8;i++{
        encrypt[i+16]=p[i]
    }
    for i:=0;i<4;i++{
        encrypt[i+24]=sig[i]
    }
    // fmt.Println("encrypt:",encrypt)

	//base64加密
    code := base64.StdEncoding.EncodeToString(encrypt)
    return code 
}


