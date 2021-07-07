package models

import "github.com/beego/beego/v2/client/orm"

type Region struct {
	Id   int
	Name string
}

type Type struct {
	Id   int
	Name string
}

func (u *Region) TableName() string {
	return "channel_region"
}

func init() {
	orm.RegisterModel(new(Region))
	orm.RegisterModel(new(Type))
}

// GetChannelRegion 获取频道下的所有地区信息
func GetChannelRegion(channelId int) (int64, []Region, error) {
	newOrm := orm.NewOrm()
	var regions []Region
	rows, err := newOrm.Raw("SELECT id,name FROM channel_region where status =1 "+
		"AND channel_id = ? ORDER BY sort DESC", channelId).QueryRows(&regions)
	return rows, regions, err
}

// GetChannelType 获取频道下的所有类型信息
func GetChannelType(channelId int) (int64, []Type, error) {
	newOrm := orm.NewOrm()
	var types []Type
	rows, err := newOrm.Raw("SELECT id,name FROM channel_type where status =1 "+
		"AND channel_id = ? ORDER BY sort DESC", channelId).QueryRows(&types)
	return rows, types, err
}
