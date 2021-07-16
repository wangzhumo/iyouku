package models

import (
	rabbitmqClient "com.wangzhumo.iyouku/services/rabbitmq"
	"encoding/json"
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

// SendMQMessageToUser 保存需要发送的消息到用户 - 这里使用RabbitMQ
func SendMQMessageToUser(messageId int64, uid int) {
	// 组织数据
	type Data struct {
		UserId    int
		MessageId int64
	}

	var data = Data{
		UserId:    uid,
		MessageId: messageId,
	}

	dataJson, _ := json.Marshal(&data)
	rabbitmqClient.Publish("", "iyouku_push_message_user", string(dataJson))
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
