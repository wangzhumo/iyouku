package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Message struct{
	Id int
	Content string
	AddTime int64
}


type MessageUser struct {
	Id int
	MessageId int
	AddTime int64
	Status int
}

func init() {
	orm.RegisterModel(new(Message))
	orm.RegisterModel(new(MessageUser))
}

// SendMessage 保存需要发送的消息
func SendMessage(content string) (int64,error)  {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now().Unix()
	count, err := o.Insert(&message)
	return count, err
}
