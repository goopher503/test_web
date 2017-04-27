package request

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/NewTrident/GoDSP/common/consts"
	"github.com/NewTrident/baidu"
	"github.com/pborman/uuid"
)

type Request struct {
	BidRequest *baidurtb.BidRequest
}

//制作完整的Request
func (r *Request) MakeBesRequest() error {
	var err error
	//对request进行随机生成
	if r.BidRequest, err = InitBesRequest(); err != nil {
		return err
	}
	//对request进行二次填充
	r.FillBesRequest()
	return nil
}

//随机生成数据填充Request
func InitBesRequest() (*baidurtb.BidRequest, error) {
	//随机种子
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i, n int
	//要添加的项
	var id, ip, user_agent, baidu_user_id, detected_language, url, referer string
	var baidu_user_id_version, site_category, site_quality, page_type, page_quality int32
	var user_category []int64
	var page_keyword []string
	var excluded_product_category []int32
	var is_test bool	
	var gender baidurtb.BidRequest_Gender
	var adslot []*baidurtb.BidRequest_AdSlot
	var mobile baidurtb.BidRequest_Mobile
	var user_geo_info baidurtb.BidRequest_Geo
	bidrequest := &baidurtb.BidRequest{}

    //id
	//32位随机字符串，0-z
	for i = 0; i < 32; i++ {
		n = seed.Intn(36)
		if n > 9 {
			id = id + string(n+87)
		} else {
			id = id + string(n+48)
		}
	}

	//ip
	//4个1-255的点分十进制字符串
	for i = 0; i < 4; i++ {
		n = seed.Intn(255) + 1
		ip = ip + strconv.Itoa(n)
		if i < 3 {
			ip = ip + "."
		}
	}
	// user_agent
	// 格式未知
	// 暂时写死
	user_agent = "Mozilla/5.0 (Linux; U; Android 2.3.6; it-it; GT-S5570I Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 (Mobile; afma-sdk-a-v6.1.0),gzip(gfe)"

	//baidu_user_id
	//32位随机字符串，0-z
	for i = 0; i < 32; i++ {
		n = seed.Intn(36)
		if n > 9 {
			baidu_user_id = baidu_user_id + string(n+87)
		} else {
			baidu_user_id = baidu_user_id + string(n+48)
		}
	}

	//baidu_user_id_version
	//int32 暂时设置为1
	baidu_user_id_version = 1

	//detected_language
	//页面语言
	//设置为"en","zh-cn","zh-tw"随机
	n = seed.Intn(3)
	if n == 0 {
		detected_language = "en"
	} else if n == 1 {
		detected_language = "zh-cn"
	} else {
		detected_language = "zh-tw"
	}

	//url
	//当前页面url
	//设置为"www.baidu+n.com",n为1-10
	n = seed.Intn(10) + 1
	url = "www.baidu" + strconv.Itoa(n) + ".com"

	//referer
	//请求的refer
	//设置为"www.baidu_refer+n.com",n为1-10
	n = seed.Intn(10) + 1
	referer = "www.baidu_refer" + strconv.Itoa(n) + ".com"

	//page_type
	//页面类型
	//设置为0-13
	page_type = seed.Int31n(14)

	//page_quality
	//页面内容质量
	//设置为0-1
	page_quality = seed.Int31n(2)

	//site_category
	//网站分类
	//设置为1-20
	site_category = seed.Int31n(21)

	//site_quality
	//网站质量类型
	//设置为0-3
	site_quality = seed.Int31n(4)

	//user_category
	//用户分类
	//设置4个，值为2+4/5/6/7+00+1-9+0
	for i = 0; i < 4; i++ {
		user_category = append(user_category, int64(200000+(seed.Intn(4)+4)*10000+(seed.Intn(9)+1)*10))
	}

	//gender
	//用户性别
	n = seed.Intn(3)
	if n == 0 {
		gender = baidurtb.BidRequest_UNKNOWN
	} else if n == 1 {
		gender = baidurtb.BidRequest_MALE
	} else {
		gender = baidurtb.BidRequest_FEMALE
	}

	//page_keyword
	//页面关键字
	//设置1-5个，值为"baidu+n"
	n = seed.Intn(5) + 1
	for i = 0; i < n; i++ {
		page_keyword = append(page_keyword, "baidu"+strconv.Itoa(seed.Intn(10)+1))
	}

	//excluded_product_category
	//发布商不允许的广告行业
	//设置1-10个，值为66-99
	n = seed.Intn(10) + 1
	for i = 0; i < n; i++ {
		excluded_product_category = append(excluded_product_category, seed.Int31n(34)+66)
	}

	//is_test
	//设置为true
	is_test = true

	//adslot
	//默认1个
	var ad baidurtb.BidRequest_AdSlot
	for i=0;i<1;i++{
	  MakeAdSlot(&ad,seed)
	  adslot = append(adslot, &ad)
	}
	//mobile
	//移动设备信息
	MakeMobile(&mobile,seed)

	//user_geo_info
	//位置信息
	MakeGeo(&user_geo_info,seed)

	bidrequest.Id = &id
	bidrequest.Ip = &ip
	bidrequest.UserAgent = &user_agent
	bidrequest.BaiduUserId = &baidu_user_id
	bidrequest.BaiduUserIdVersion = &baidu_user_id_version
	bidrequest.DetectedLanguage = &detected_language
	bidrequest.Url = &url
	bidrequest.Referer = &referer
	bidrequest.PageType = &page_type
	bidrequest.PageQuality = &page_quality
	bidrequest.SiteCategory = &site_category
	bidrequest.SiteQuality = &site_quality
	bidrequest.UserCategory = user_category
	bidrequest.Gender = &gender
	bidrequest.PageKeyword = page_keyword
	bidrequest.ExcludedProductCategory = excluded_product_category
	bidrequest.IsTest = &is_test
	bidrequest.Adslot = adslot
	bidrequest.Mobile = &mobile
	bidrequest.UserGeoInfo = &user_geo_info

	return bidrequest, nil
}

func MakeAdSlot(a *baidurtb.BidRequest_AdSlot,seed *rand.Rand) {

	var n, i int

	//需要添加的数据
	var ad_block_key uint64
	var adslot_level baidurtb.BidRequest_AdSlot_AdSlotLevel
	var sequence_id, adslot_type, width, height, slot_visibility, minimum_cpm int32
	var creative_type []int32
	var creative_desc_type []baidurtb.BidRequest_AdSlot_CreativeDescType
	var excluded_landing_page_url []string

	//ad_block_key
	//广告位id
	//100000以内随机数
	ad_block_key = uint64(seed.Intn(100000))

	//sequence_id
	//广告位顺序id
	//1-20
	sequence_id = seed.Int31n(20) + 1

	//adslot_level
	//广告等级0,1,2,3,4
	n=seed.Intn(5)
	switch n {
	case 0:
		adslot_level=baidurtb.BidRequest_AdSlot_UNKNOWN_ADB_LEVEL
	case 1:
		adslot_level=baidurtb.BidRequest_AdSlot_TOP
	case 2:
		adslot_level=baidurtb.BidRequest_AdSlot_MED
	case 3:
		adslot_level=baidurtb.BidRequest_AdSlot_TAIL
	case 4:
		adslot_level=baidurtb.BidRequest_AdSlot_LOW
	}
	


	//adslot_type
	//展示类型
	//0|1|11|12
	n = seed.Intn(4)
	switch n {
	case 0:
		adslot_type = 0
	case 1:
		adslot_type = 1
	case 2:
		adslot_type = 11
	case 3:
		adslot_type = 12

	}

	//w 和 h
	//从const中随机选取
	n = seed.Intn(2)
	switch n {
	case 0:
		height = consts.BannerRectangleHight
		width = consts.BannerRectangleWidth
	case 1:
		height = consts.BannerRegularHight
		width= consts.BannerRegularWidth
	}


	//slot_visibility
	//展示位置
	//0|1
	slot_visibility = seed.Int31n(2)

	//minimum_cpm
	//底价
	//1-1000
	minimum_cpm = seed.Int31n(100) + 1

	//creative_type
	//发布商允许的类型
	//1个，值为1
	creative_type = append(creative_type, 1)

	//creative_desctype
	//创意封装与渲染
	//两个值
	var one_creative_desc_type baidurtb.BidRequest_AdSlot_CreativeDescType
	for i = 0; i < 1; i++ {
		n = seed.Intn(2)
		switch n + 1 {
		case 1:
			one_creative_desc_type = baidurtb.BidRequest_AdSlot_STATIC_CREATIVE
		case 2:
			one_creative_desc_type = baidurtb.BidRequest_AdSlot_DYNAMIC_CREATIVE
		}
		creative_desc_type = append(creative_desc_type, one_creative_desc_type)
	}

	//excluded_landing_page_url
	//发布商不允许的url
	//1-10个，值为"www.landing_page"+n+".com",n为1-100
	n = seed.Intn(10) + 1
	for i = 0; i < n; i++ {
		excluded_landing_page_url = append(excluded_landing_page_url, "www.landing_page"+strconv.Itoa(seed.Intn(100)+1)+".com")
	}

	a.AdBlockKey = &ad_block_key
	a.SequenceId = &sequence_id
	a.AdslotType = &adslot_type
	a.AdslotLevel= &adslot_level
	a.Width = &width
	a.Height = &height
	a.SlotVisibility = &slot_visibility
	a.MinimumCpm = &minimum_cpm
	a.CreativeType = creative_type
	a.CreativeDescType = creative_desc_type
	a.ExcludedLandingPageUrl = excluded_landing_page_url
}

func MakeMobile(m *baidurtb.BidRequest_Mobile,seed *rand.Rand) {

	var n,i int

	var device_type baidurtb.BidRequest_Mobile_MobileDeviceType
	var platform baidurtb.BidRequest_Mobile_OS
	var os_version baidurtb.BidRequest_Mobile_DeviceOsVersion
	var screen_width, screen_height int32
	var carrier_id int64
	var mobile_app baidurtb.BidRequest_Mobile_MobileApp
	var for_advertising_id []*baidurtb.BidRequest_Mobile_ForAdvertisingID

	//device_type
	//设备类型
	n = seed.Intn(3)
	switch n {
	case 0:
		device_type = baidurtb.BidRequest_Mobile_UNKNOWN_DEVICE
	case 1:
		device_type = baidurtb.BidRequest_Mobile_HIGHEND_PHONE
	case 2:
		device_type = baidurtb.BidRequest_Mobile_TABLET
	}

	//platform
	//设备平台
	n = seed.Intn(2)
	switch n {
	case 0:
		platform = baidurtb.BidRequest_Mobile_ANDROID
	case 1:
		platform = baidurtb.BidRequest_Mobile_IOS
	}

	//for_advertising_id
	//为了广告业的id
	var a_for_advertising_id baidurtb.BidRequest_Mobile_ForAdvertisingID
    for i=0;i<1;i++{
		MakeForAdvertisingID(&a_for_advertising_id, n+1)
		for_advertising_id = append(for_advertising_id, &a_for_advertising_id)
	}
	//os_version
	//系统版本
	//三个都是0-9随机
	major := seed.Int31n(10)
	minor := seed.Int31n(10)
	micro := seed.Int31n(10)
	os_version.OsVersionMajor = &major
	os_version.OsVersionMinor = &minor
	os_version.OsVersionMicro = &micro

	//screen_height 和 screen_width
	//设备屏幕宽高
	//暂时写死
	screen_height = 1280
	screen_width = 960

	//carrier_id
	//运营商编号
	//暂时写死
	carrier_id = 46000

	//mobile_app
	//移动应用信息
	MakeMobileApp(&mobile_app,seed)

	m.DeviceType = &device_type
	m.Platform = &platform
	m.OsVersion = &os_version
	m.ScreenWidth = &screen_width
	m.ScreenHeight = &screen_height
	m.CarrierId = &carrier_id
	m.MobileApp = &mobile_app
	m.ForAdvertisingId = for_advertising_id
}

func MakeMobileApp(ma *baidurtb.BidRequest_Mobile_MobileApp,seed *rand.Rand) {

	var n, i int

	var app_id, app_bundle_id string
	var app_category int32
	var app_interaction_type []baidurtb.BidRequest_Mobile_MobileApp_AppInteractionType

	//app_id
	//8位0-z
	for i = 0; i < 8; i++ {
		n = seed.Intn(36)
		if n > 9 {
			app_id = app_id + string(n+87)
		} else {
			app_id = app_id + string(n+48)
		}
	}

	//app_bundle_id
	// 如果来自苹果商店，则直接是app-store id
	// 如果来自Android设备，则是package的全名
	//"test_app_id"+n,n为0-9
	app_bundle_id = "test_app_id" + strconv.Itoa(seed.Intn(10))

	//app_category
	//App应用分类
	//5000-9999
	app_category = seed.Int31n(5000) + 5000

	//app_interaction_type
	// App允许的交互类型定义
	// 电话、下载、应用唤醒
	//暂时默认为1个
	n = seed.Intn(3)
	switch n {
	case 0:
		app_interaction_type = append(app_interaction_type, baidurtb.BidRequest_Mobile_MobileApp_TELEPHONE)
	case 1:
		app_interaction_type = append(app_interaction_type, baidurtb.BidRequest_Mobile_MobileApp_DOWNLOAD)
	case 2:
		app_interaction_type = append(app_interaction_type, baidurtb.BidRequest_Mobile_MobileApp_DEEPLINK)

	}

	ma.AppId = &app_id
	ma.AppBundleId = &app_bundle_id
	ma.AppCategory = &app_category
	ma.AppInteractionType = app_interaction_type
}

func MakeGeo(g *baidurtb.BidRequest_Geo,seed *rand.Rand) {

	var n int

	//var user_coordinate []*baidurtb.BidRequest_Geo_Coordinate
	var user_location baidurtb.BidRequest_Geo_UserLocation

	//user_location
	//用户位置信息
	var province, city string

	n = seed.Intn(3)
	switch n {
	case 0:
		province = "北京"
		city = "北京"
	case 1:
		province = "上海"
		city = "上海"
	case 2:
		province = "四川"
		city = "成都"
	}
	user_location.Province = &province
	user_location.City = &city
	//g.UserCoordinate = user_coordinate
	g.UserLocation = &user_location
}

func MakeForAdvertisingID(f *baidurtb.BidRequest_Mobile_ForAdvertisingID, t int) {

	var Type baidurtb.BidRequest_Mobile_ForAdvertisingID_IDType
	var Id string
	switch t {
	case 0:
		Type = baidurtb.BidRequest_Mobile_ForAdvertisingID_UNKNOWN
	case 1:
		Type = baidurtb.BidRequest_Mobile_ForAdvertisingID_ANDROID_ID
	case 2:
		Type = baidurtb.BidRequest_Mobile_ForAdvertisingID_IDFA
	}
	Id = uuid.NewRandom().String()
	f.Type = &Type
	f.Id = &Id
}
