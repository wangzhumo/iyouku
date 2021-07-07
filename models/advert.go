package models

import "github.com/beego/beego/v2/client/orm"

type Advert struct {
	Id       int
	Title    string
	SubTitle string
	AddTime  int64
	Img      string
	Url      string
}

func init() {
	orm.RegisterModel(new(Advert))
}

func (u *Advert) TableName() string {
	return "advert"
}

// GetAdvertByChannelID 通过ChannelID获取Advert数据
func GetAdvertByChannelID(channelId int) (int64, []Advert, error) {
	newOrm := orm.NewOrm()
	var adverts []Advert
	count, err := newOrm.Raw("SELECT id, title, sub_title, channel_id, add_time,img,url FROM advert where status = 1 AND channel_id =? ORDER BY sort DESC LIMIT 1", channelId).QueryRows(&adverts)
	return count, adverts, err
}
