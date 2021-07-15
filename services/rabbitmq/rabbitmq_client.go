package rabbitmqClient

import (
	"bytes"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/streadway/amqp"
)

type MQCallback func(message string)

// ConnectMq 连接RabbitMQ
func ConnectMq() (conn *amqp.Connection, err error) {
	rabbitMq, err := beego.AppConfig.String("rabbitmqdb")
	conn, err = amqp.Dial(rabbitMq)
	return
}

// Publish 发送一个消息
func Publish(exchange string, queueName string, message string) error {
	// 连接
	conn, err := ConnectMq()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 发送消息即可
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	// 创建队列
	//	Queue:name,
	//	Durable:    durable,
	//	AutoDelete: autoDelete,
	//	Exclusive:  exclusive,
	//	NoWait:     noWait,
	//	Arguments:  args,
	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// 发送
	err = channel.Publish(exchange, queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	return err
}

// Consumer 接收者
func Consumer(exchange string, queueName string, callback MQCallback) {
	// 连接
	conn, err := ConnectMq()
	defer conn.Close()
	if err != nil {
		return
	}

	//创建通道
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		return
	}
	// 创建队列
	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	messages, err := channel.Consume(queue.Name, "", true,
		false, false, false, nil)

	if err != nil {
		return
	}

	// 监听消息，需要一直监听，所以不停止
	forever := make(chan bool)
	// 使用协程
	go func() {
		for message := range messages {
			messageStr := BytesToString(&(message.Body))
			// 转发给外部
			callback(*messageStr)
			message.Ack(false)
		}
	}()
	<-forever
}

// BytesToString 字节转string
func BytesToString(b *[]byte) *string {
	buffer := bytes.NewBuffer(*b)
	s := buffer.String()
	return &s
}
