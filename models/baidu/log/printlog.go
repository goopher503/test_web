package printlog

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/NewTrident/baidu"
)

func BaiDuPrintLog(request *baidurtb.BidRequest, response *baidurtb.BidResponse, num int) {
	err := os.MkdirAll("models/baidu/log/log", 0777)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	var filename string

	filename = "models/baidu/log/log/" + tm.Format("2006_01_02_15_04_05") + ".log"

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0664) //创建文件

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "", 0)
	logger.Println(num + 1)
	PrintRequest(request, logger)
	logger.Println("\r\n\r\n\r\n\r\n")
	PrintResponse(response, logger)
	logger.Println("\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n")
}

func PrintRequest(r *baidurtb.BidRequest, l *log.Logger) {
	l.Println("=== BidRequest: ===")
	l.Println("id: ", r.GetId())
	l.Println("ip: ", r.GetIp())
	l.Println("user_agent: ", r.GetUserAgent())
	l.Println("baidu_user_id: ", r.GetBaiduUserId())
	l.Println("baidu_user_id_version: ", r.GetBaiduUserIdVersion())
	for i := range r.GetUserCategory() {
		l.Println("user_category: ", r.GetUserCategory()[i])
	}
	l.Println("gender: ", r.GetGender())
	l.Println("detected_language: ", r.GetDetectedLanguage())
	l.Println("url: ", r.GetUrl())
	l.Println("referer: ", r.GetReferer())
	l.Println("site_category: ", r.GetSiteCategory())
	l.Println("site_quality: ", r.GetSiteQuality())
	l.Println("page_type: ", r.GetPageType())
	for i := range r.GetPageKeyword() {
		l.Println("page_keyword: ", r.GetPageKeyword()[i])
	}

	l.Println("page_quality: ", r.GetPageQuality())
	for i := range r.GetExcludedProductCategory() {
		l.Println("excluded_product_category: ", r.GetExcludedProductCategory()[i])
	}
	for i := range r.GetAdslot() {
		l.Println("adslot { ")
		l.Println("	ad_block_key: ", r.GetAdslot()[i].GetAdBlockKey())
		l.Println("	sequence_id: ", r.GetAdslot()[i].GetSequenceId())
		l.Println("	adslot_type: ", r.GetAdslot()[i].GetAdslotType())
		l.Println("	width: ", r.GetAdslot()[i].GetWidth())
		l.Println("	height: ", r.GetAdslot()[i].GetHeight())
		l.Println("	slot_visibility: ", r.GetAdslot()[i].GetSlotVisibility())
		for j := range r.GetAdslot()[i].GetCreativeType() {
			l.Println("	creative_type: ", r.GetAdslot()[i].GetCreativeType()[j])
		}
		for j := range r.GetAdslot()[i].GetExcludedLandingPageUrl() {
			l.Println("	excluded_landing_page_url: ", r.GetAdslot()[i].GetExcludedLandingPageUrl()[j])
		}
		l.Println("	minimum_cpm: ", r.GetAdslot()[i].GetMinimumCpm())
		l.Println("} ")
	}

	l.Println("is_test: ", r.GetIsTest())
	l.Println("user_geo_info { ")
	l.Println("  user_location { ")
	l.Println("    province: ", r.GetUserGeoInfo().GetUserLocation().GetProvince())
	l.Println("    city: ", r.GetUserGeoInfo().GetUserLocation().GetCity())
	l.Println("  }\r\n} ")
	l.Println("mobile { ")
	l.Println("  device_type: ", r.GetMobile().GetDeviceType())
	l.Println("  platform: ", r.GetMobile().GetPlatform())
	l.Println("  os_version {")
	l.Println("    os_version_major: ", r.GetMobile().GetOsVersion().GetOsVersionMajor())
	l.Println("    os_version_minor: ", r.GetMobile().GetOsVersion().GetOsVersionMinor())
	l.Println("    os_version_micro: ", r.GetMobile().GetOsVersion().GetOsVersionMicro())
	l.Println("  } ")
	for i := range r.GetMobile().GetForAdvertisingId() {
		l.Println("  for_advertising_id { ")
		l.Println("    type :", r.GetMobile().GetForAdvertisingId()[i].GetType())
		l.Println("    id :", r.GetMobile().GetForAdvertisingId()[i].GetId())
		l.Println("  }")
	}
	l.Println("  screen_width: ", r.GetMobile().GetScreenWidth())
	l.Println("  screen_height: ", r.GetMobile().GetScreenHeight())
	l.Println("  carrier_id: ", r.GetMobile().GetCarrierId())
	l.Println("}")

}

func PrintResponse(r *baidurtb.BidResponse, l *log.Logger) {
	l.Println("=== BidResponse: ===")
	l.Println("id: ", r.GetId())
	for i := range r.GetAd() {
		l.Println("ad {")
		l.Println("  sequence_id: ", r.GetAd()[i].GetSequenceId())
		l.Println("  creative_id: ", r.GetAd()[i].GetCreativeId())
		l.Println("  adm:",r.GetAd()[i].GetHtmlSnippet())
		l.Println("  max_cpm: ", r.GetAd()[i].GetMaxCpm())
		l.Println("  is_cookie_matching: ", r.GetAd()[i].GetIsCookieMatching())
		l.Println("}")
	}
	l.Println("processing_time_ms: ", r.GetProcessingTimeMs())
}
