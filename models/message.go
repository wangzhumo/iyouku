package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	MessageId int64
	AddTime   int64
	Status    int
	UserId    int
}

func init() {
	orm.RegisterModel(new(Message))
	orm.RegisterModel(new(MessageUser))
}

// SendMessage 保存需要发送的消息
func SendMessage(content string) (int64, error) {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Now().Unix()
	messageId, err := o.Insert(&message)
	return messageId, err
}

// SendMessageToUser 保存需要发送的消息到用户
func SendMessageToUser(messageId int64, uid int) (int64, error) {
	o := orm.NewOrm()
	var message MessageUser
	message.MessageId = messageId
	message.AddTime = time.Now().Unix()
	message.Status = 1
	message.UserId = uid
	messageUserId, err := o.Insert(&message)
	return messageUserId, err
}
