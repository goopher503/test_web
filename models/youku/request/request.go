package request

import (
	"math/rand"
	"time"
	"strconv"
	"github.com/pborman/uuid"

	"github.com/NewTrident/GoDSP/common/consts"
	"github.com/NewTrident/youku"
)

type Request struct {
	BidRequest *youku.BidRequest
}

//制作完整的Request
func (r *Request) MakeYkRequest() error {
	var err error
	if r.BidRequest, err = InitYkRequest(); err != nil {
		return err
	}
	r.FillYkRequest()
	return nil
}

//随机生成数据填充Request
func InitYkRequest() (*youku.BidRequest, error) {

	//随机种子
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i, n int

	//要添加的项
	var id string
	var imp []youku.Imp
	var site youku.Site
	var app youku.App
	var device youku.Device
	var user youku.User
	bidrequest := &youku.BidRequest{}

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

	//imp
	//随机发送三条以内的imp
	
	for i = 0; i < 3; i++ {
		var ad youku.Imp
		
		MakeImp(&ad,seed)
		imp = append(imp, ad)
	}

	//app与site
	n = seed.Intn(1)
	if n == 0 {
		MakeApp(&app)
		bidrequest.App = &app
	}
	if n == 1 {
		MakeSite(&site)
		bidrequest.Site = &site
	}

	//device
	MakeDevice(&device,seed)

	//user
	MakeUser(&user,seed)

	bidrequest.ID=id
	bidrequest.Imp=imp
	bidrequest.Site=&site
	bidrequest.App=&app
	bidrequest.Device=&device
	bidrequest.User=&user

	return bidrequest,nil

}
func MakeImp(im *youku.Imp,seed *rand.Rand) {

	var i, n int

	//要添加的项
	var id, tagid string
	var bidfloor float64
	var banner youku.Banner
	// var video  *youku.Video
	// var native *youku.Native
	var pmp youku.Pmp
	var ext youku.ImpExt
	var secure int8

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

	//tagid
	//10位随机字符串，0-z
	for i = 0; i < 10; i++ {
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

	//banner
	MakeBanner(&banner,seed)

	//pmp
	MakePmp(&pmp,seed)

	//ext
	ext.Repeat = seed.Int63n(10)

	//secure
	secure = int8(seed.Int63n(10))

	im.ID = id
	im.TagID = tagid
	im.BidFloor = bidfloor
	im.Banner = &banner
	im.Pmp = &pmp
	im.Ext = &ext
	im.Secure = secure

}
func MakeBanner(b *youku.Banner,seed *rand.Rand) {

	var n int

	//要添加的项
	var width, height uint64
	var pos int8

	//w 和 h
	//从const中随机选取
	n = seed.Intn(2)
	switch n {
	case 0:
		height = consts.BannerRectangleHight
		width = consts.BannerRectangleWidth
	case 1:
		height = consts.BannerRegularHight
		width = consts.BannerRegularWidth
	}

	//pos
	//不验证，暂时写死
	pos = 1

	b.W = width
	b.H = height
	b.Pos = pos
}
func MakePmp(p *youku.Pmp,seed *rand.Rand) {

	var  i int

	//要添加的项
	var id string
	var at int8
	var bidfloor float64

	//id
	id = string(seed.Intn(10) + 48)

	//at
	at = int8(seed.Intn(2))

	//bidfloor
	bidfloor = seed.Float64()

	for i = 1; i < 1; i++ {
		p.Deals[i].ID = id
		p.Deals[i].AT = at
		p.Deals[i].BidFloor = bidfloor

	}

}

// func MakeExt(e *youku.ImpExt) {

// 	//随机种子
// 	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	e.Repeat = seed.Int63n(10)
// }
func MakeApp(a *youku.App) {

	//要添加的项
	var name string
	var content youku.Content

	//name
	name="优酷客户端"

	//content
	MakeContent(&content)

	a.Content=&content
	a.Name=name

}
func MakeContent(c *youku.Content){

	var title,keywords string
    var ext youku.ContentExt
	
	//暂时写死
	title = "Big Bang Theory"
	keywords="Sitcom|Funny"
	MakeContentExt(&ext)

	c.Ext=&ext
	c.KeyWords=keywords
	c.Title=title

}
func MakeContentExt(e *youku.ContentExt){

	var channel,cs,usr,s,vid string
    
	//暂时写死
	 channel= "c"
     usr= "97454045"
     cs="2034"
     vid= "420613008"
     s= "306281"

	 e.Channel=channel
	 e.CS=cs
	 e.S=s
	 e.Usr=usr
	 e.Vid=vid
}

func MakeSite(s *youku.Site){

	var name,page,ref string
	var content youku.Content

	//暂时写死
	name="优酷网"
	page="http://www.youku.com/"
	ref=""
	MakeContent(&content)

	s.Content=&content
	s.Name=name
	s.Page=page
	s.Ref=ref

}

func MakeDevice(d *youku.Device,seed *rand.Rand){

	var n, i int

	//要添加的项
	var ip,ua,os,osv,idfa,androidid,make,model string
	var devicetype int8

	//ip
	//4个1-255的点分十进制字符串
	for i = 0; i < 4; i++ {
		n = seed.Intn(255) + 1
		ip = ip + strconv.Itoa(n)
		if i < 3 {
			ip = ip + "."
		}
	}

	//ua
	//不验证，暂时写死
	ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3"

    // os
	n = seed.Intn(2)
	switch n {
	case 0:
		os = "iOS"
		idfa=uuid.NewRandom().String()
		androidid=""
	case 1:
		os = "Android"
		androidid=uuid.NewRandom().String()
		idfa=""
	}
	// os="iOS"

	//osv
	osv = string(seed.Intn(10)+48) + "." + string(seed.Intn(10)+48)
	// osv="8.0"

	//make
	//不验证，暂时写死
	make = "Apple"

	//model
	//不验证，暂时写死
	model = "iPhone"

	//mac地址

	//devicetype
	//0—手机，1—平板，2—PC
	n=seed.Intn(3)
	devicetype=int8(n)

	d.AndroidID=androidid
	d.DeviceType=devicetype
	d.IDFA=idfa
	d.IP=ip
	d.Make=make
	d.Model=model
	d.OS=os
	d.OSV=osv
	d.UA=ua

}
func MakeUser(u *youku.User,seed *rand.Rand){

	var n, i int

	var id,gender string

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

	//gender
	gender ="M"	

	u.ID=id
	u.Gender=gender
}