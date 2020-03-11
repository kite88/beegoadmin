package models

import (
	"beegoadmin/utils"
	"beegoadmin/utils/page"
	"time"
)

type Message struct {
	Id         uint
	Title      string
	Content    string
	AdminId    uint
	IsRead     int
	ReadTime   int64
	CreateTime int64
	Type       int

	CreateTimeFormat string `gorm:"-" json:"CreateTime"`
}

//获取数量
func GetMessageCount(query interface{}, args ...interface{}) (count int) {
	Db.Model(&Message{}).Where(query, args...).Count(&count)
	return
}

//获取单条消息
func FindMessageByMessage(message Message) (val Message, err error) {
	err = Db.Where(message).First(&val).Error
	var format = "2006/01/02 15:04"
	if time.Unix(val.CreateTime, 0).Year() == time.Now().Year() {
		format = "01/02 15:04"
		if time.Unix(val.CreateTime, 0).Month() == time.Now().Month() && time.Unix(val.CreateTime, 0).Day() == time.Now().Day() {
			format = "15:04"
		}
	}
	val.CreateTimeFormat = utils.FormatDateTime(val.CreateTime, format)
	return
}

//获取数据
func GetMessageData(viewRows uint, p uint, query interface{}, args ...interface{}) page.Data {
	var count uint
	var list []Message

	Db.Where(query, args...).Find(&list).Count(&count)
	data, limit, offset := page.Page(count, viewRows, p)
	if len(list) > 0 {
		Db.Where(query, args...).
			Limit(limit).
			Offset(offset).
			Order("create_time desc").
			Find(&list)
		for key, val := range list {
			format := "15:04 2006/01/02"
			if time.Unix(val.CreateTime, 0).Year() == time.Now().Year() {
				format = "15:04 01/02"
			}
			list[key].CreateTimeFormat = utils.FormatDateTime(val.CreateTime, format)
		}
	}
	data.List = list
	return data
}

//新增数据
func AddMessage(value Message) error {
	value.CreateTime = time.Now().Unix()
	err := Db.Create(value).Error
	return err
}

//更新数据
func UpdateMessage(value map[string]interface{}, query interface{}, args ...interface{}) error {
	err := Db.Model(&Message{}).Where(query, args...).Updates(value).Error
	return err
}

func (m *Message) TableName() string {
	return TableName("message")
}
