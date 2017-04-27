package request

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/NewTrident/GoDSP/common/consts"
	"github.com/NewTrident/openrtb"
	"github.com/pborman/uuid"
)

type Request struct {
	BidRequest *openrtb.BidRequest
}

//制作完整的Request
func (r *Request) MakeRequest() error {
	var err error
	if r.BidRequest, err = InitRequest(); err != nil {
		return err
	}
	r.FillRequest()
	return nil
}

//随机生成数据填充Request
func InitRequest() (*openrtb.BidRequest, error) {

	//随机种子
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i, n int

	//要添加的项
	var id string
	var at int8
	// var test, allimps int8
	var tmax uint64
	var imp []openrtb.Imp
	var app openrtb.App
	var site openrtb.Site
	var device openrtb.Device
	var user openrtb.User
	// var wseat, cur []string
	var bcat, badv []string
	var ext openrtb.SmaatoExt
	// var reg *openrtb.Regs
	bidrequest := &openrtb.BidRequest{}

	//id
	//10位随机字符串，0-z
	for i = 0; i < 10; i++ {
		n = seed.Intn(36)
		if n > 9 {
			id = id + string(n+87)
		} else {
			id = id + string(n+48)
		}
	}

	//at
	//拍卖类型：1或者2
	n = seed.Intn(2)
	if n == 0 {
		at = 1
	}
	if n == 1 {
		at = 2
	}

	// test
	// 是否为测试报文
	// test = int8(seed.Int31n(2))

	// allimps
	// ...
	// allimps = int8(seed.Intn(2))

	//tmax
	//请求超时设置(ms)
	tmax = 300

	//imp
	//默认1个
	var ad openrtb.Imp
	for i = 0; i < 1; i++ {
		MakeImp(&ad, i,seed)
		imp = append(imp, ad)
	}

	//app与site
	n = seed.Intn(1)
	if n == 0 {
		MakeApp(&app,seed)
		bidrequest.App = &app
	}
	if n == 1 {
		MakeSite(&site,seed)
		bidrequest.Site = &site
	}

	//device
	MakeDevice(&device, n,seed)

	//user
	MakeUser(&user,seed)

	//bcat, badv
	//不验证，暂时写死
	bcat = []string{
		"IAB25",
		"IAB7-39",
		"IAB8-18",
		"IAB8-5",
		"IAB9-9",
	}
	badv = []string{
		"apple.com",
		"go-text.me",
		"heywire.com",
	}

	//ext
	MakeExt(&ext, &device)
	bidrequest.ID = id
	bidrequest.AT = at
	bidrequest.TMax = tmax
	bidrequest.Imp = imp
	bidrequest.Device = &device
	bidrequest.User = &user
	bidrequest.BCat = bcat
	bidrequest.BAdv = badv
	bidrequest.Ext = &ext
	return bidrequest, nil
}
func MakeImp(a *openrtb.Imp, idd int,seed *rand.Rand) {

	var n, i int

	//需要添加的数据
	var id, tagid string
	var bidfloor float64
	// var bidFloorCur string
	var banner openrtb.Banner
	// var video *openrtb.Video
	// var native *openrtb.Native
	// var ext *openrtb.ImpExt

	//id
	id = string(idd + 48)

	//tagid
	//10位随机字符串，0-z
	for i = 0; i < 35; i++ {
		n = seed.Intn(36)
		if n > 9 {
			id = id + string(n+87)
		} else {
			id = id + string(n+48)
		}
	}

	//bidfloor
	//广告位最低价
	bidfloor = seed.Float64()

	//banner,video,native选择
	// n = seed.Intn(3)
	// if n == 0 {
	// 	var ban *openrtb.Banner
	// 	MakeBanner(ban)
	// 	a.Banner = ban
	// }
	// if n == 1 {
	// 	var vi *openrtb.Video
	// 	MakeVideo(vi)
	// 	a.Video = vi

	// }
	// if n == 2 {
	// 	var na *openrtb.Native
	// 	MakeNative(na)
	// 	a.Native = na

	// }

	//ext
	// ext.Strictbannersize = 0

	//banner
	MakeBanner(&banner,seed)

	a.Banner = &banner
	a.ID = id
	a.BidFloor = bidfloor
	// a.Ext = ext
	a.TagID = tagid
}
func MakeBanner(b *openrtb.Banner,seed *rand.Rand) {

	var n,i int

	//需要填充的项
	var id string
	var w, h uint64
	var pos int8
	var btype, battr, api []int8
	var format []openrtb.Format
	var mimes []string

	//id
	//不知道怎么填，暂时写成10位字符串
	for i = 0; i < 10; i++ {
		n = seed.Intn(36)
		if n > 9 {
			id = id + string(n+87)
		} else {
			id = id + string(n+48)
		}
	}

	//w 和 h
	//从const中随机选取
	n = seed.Intn(2)
	switch n {
	case 0:
		h = consts.BannerRectangleHight
		w = consts.BannerRectangleWidth
	case 1:
		h = consts.BannerRegularHight
		w = consts.BannerRegularWidth
	}

	//pos
	//广告在屏幕上的位置
	//不验证，暂时写死
	pos = 1

	//btype
	//限制的banner类型
	//不验证，暂时写死
	btype = []int8{1}

	//battr
	//限制的物料属性
	//不验证，暂时写死
	battr = []int8{1, 2, 3, 5, 6}

	//api
	//API框架表
	//不验证，暂时写死
	api = []int8{3, 5}

	//format
	//等于w 与 h
	var form openrtb.Format
	form.W = w
	form.H = h
	format = append(format, form)

	//mimes
	//不验证，暂时写死
	mimes = []string{
		"text/javascript",
		"application/javascript",
		"image/jpeg",
		"image/png",
		"image/gif",
	}

	b.ID = id
	b.W = w
	b.H = h
	b.Format = format
	b.Pos = pos
	b.BAttr = battr
	b.BType = btype
	b.API = api
	b.MIMEs = mimes
}

// func MakeVideo(b *openrtb.Video) {

// }
// func MakeNative(b *openrtb.Native) {

// }
func MakeApp(a *openrtb.App,seed *rand.Rand) {

	var n, i int

	//需要填充的项
	var id, name, domain string
	var cat []string
	var publisher openrtb.Publisher

	//id
	//9位随数字字符串
	for i = 0; i < 9; i++ {
		n = seed.Intn(10)
		id = id + string(n+48)
	}

	//name
	//不验证，暂时写死
	name = "Aerserv TextNow iOS_AR TextNow iOS 320x50 Premium Tier_iOS_XXLARGE_320x50_IAB9"

	//domain
	//不验证，暂时写死
	domain = "itunes.apple.com"

	//cat
	//IAB内容类型数组
	//不验证，暂时写死
	cat = []string{"IAB9"}

	// //keywords
	// //不验证，暂时写死
	// keywords = ""

	//Publisher
	//不验证，暂时写死
	MakeAppPublisher(&publisher,seed)

	a.ID = id
	a.Name = name
	a.Domain = domain
	// a.Keywords = keywords
	a.Cat = cat
	a.Publisher = &publisher
}
func MakeAppPublisher(p *openrtb.Publisher,seed *rand.Rand) {

	var n, i int

	var id, name string
	//id
	//9位随数字字符串
	for i = 0; i < 9; i++ {
		n = seed.Intn(10)
		id = id + string(n+48)
	}
	name = "OpenX"

	p.ID = id
	p.Name = name
}
func MakeSite(s *openrtb.Site,seed *rand.Rand) {

	var n, i int

	//需要填充的项
	var id, name, domain string
	var cat []string
	var publisher openrtb.Publisher

	//id
	//9位随数字字符串
	for i = 0; i < 9; i++ {
		n = seed.Intn(10)
		id = id + string(n+48)
	}

	//name
	//不验证，暂时写死
	name = "Aerserv TextNow iOS_AR TextNow iOS 320x50 Premium Tier_iOS_XXLARGE_320x50_IAB9"

	//domain
	//不验证，暂时写死
	domain = "itunes.apple.com"

	//cat
	//不验证，暂时写死
	cat = []string{"IAB9"}

	// //keywords
	// //不验证，暂时写死
	// keywords = ""

	//Publisher
	//不验证，暂时写死
	MakeSitePublisher(&publisher,seed)
	s.ID = id
	s.Name = name
	s.Domain = domain
	// s.Keywords = keywords
	s.Cat = cat
	s.Publisher = &publisher

}
func MakeSitePublisher(p *openrtb.Publisher,seed *rand.Rand) {

	var n, i int
	
	var id, name string
	//id
	//9位随数字字符串
	for i = 0; i < 9; i++ {
		n = seed.Intn(10)
		id = id + string(n+48)
	}
	name = "OpenX"

	p.ID = id
	p.Name = name
}

func MakeDevice(d *openrtb.Device, as int,seed *rand.Rand) {

	var n, i int

	//需要填充的项
	var ua, ip, make, model, os, osv string
	var js int8
	var geo openrtb.Geo

	//ua
	//不验证，暂时写死
	ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3"

	//ip
	//4个1-255的点分十进制字符串
	for i = 0; i < 4; i++ {
		n = seed.Intn(255) + 1
		ip = ip + strconv.Itoa(n)
		if i < 3 {
			ip = ip + "."
		}
	}

	//make
	//不验证，暂时写死
	make = "Apple"

	//model
	//不验证，暂时写死
	model = "iPhone"

	//os
	n = seed.Intn(2)
	switch n {
	case 0:
		os = "iOS"
	case 1:
		os = "Android"
	}

	//osv
	osv = string(seed.Intn(10)+48) + "." + string(seed.Intn(10)+48)

	//js
	js = 1

	//geo
	MakeGeo(&geo)

	d.UA = ua
	d.IP = ip
	d.Make = make
	d.Model = model
	d.OS = os
	d.OSV = osv
	d.JS = js
	d.Geo = &geo
}
func MakeGeo(g *openrtb.Geo) {

	//需要填充的项
	var country, region, city string

	//不验证，暂时写死
	country = "China"
	region = "sichuan"
	city = "chengdu"

	g.Country = country
	g.Region = region
	g.City = city
}
func MakeUser(u *openrtb.User,seed *rand.Rand) {

	var n, i int

	//需要填充的项
	var id, gender string

	//id
	//9位随数字字符串
	for i = 0; i < 9; i++ {
		n = seed.Intn(10)
		id = id + string(n+48)
	}

	//gender
	gender = "F"

	u.ID = id
	u.Gender = gender
}
func MakeExt(e *openrtb.SmaatoExt, d *openrtb.Device) {
	
	var uid openrtb.UniqueDevice
	MakeUid(&uid, d)
	e.Udi = &uid
}
func MakeUid(u *openrtb.UniqueDevice, d *openrtb.Device) {

	var idfa, androidid string
	if d.OS == "iOS" {
		idfa = uuid.NewRandom().String()
		androidid = ""
	}
	if d.OS == "Android" {
		androidid = uuid.NewRandom().String()
		idfa = ""
	}
	u.IDFA = idfa
	u.GOOGLEADID = androidid
}
