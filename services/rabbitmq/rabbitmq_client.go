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
	defer conn.Close()
	if err != nil {
		return err
	}

	// 发送消息即可
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		return err
	}

	// 创建队列
	//	Queue:name,
	//	Durable:    durable,
	//	AutoDelete: autoDelete,
	//	Exclusive:  exclusive,
	//	NoWait:     noWait,
	//	Arguments:  args,
	queue, err := channel.QueueDeclare(
		queueName,
		true,
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
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(message),
	})
	return err
}

// PublishExchange 发送一个消息[添加exchange的模式]
func PublishExchange(exchange string, types string, routingKey string, message string) error {
	// 连接
	conn, err := ConnectMq()
	defer conn.Close()
	if err != nil {
		return err
	}

	// 发送消息即可
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		return err
	}

	// 创建交换机
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// 发送
	err = channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(message),
	})
	return err
}

// PublishDlx 发送一个消息[死信队列]
func PublishDlx(exchange string, message string) error {
	// 连接
	conn, err := ConnectMq()
	defer conn.Close()
	if err != nil {
		return err
	}

	// 创建Channel
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		return err
	}

	// 发送
	err = channel.Publish(exchange, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(message),
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
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	messages, err := channel.Consume(queue.Name, "", false,
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
			// 手动应答 - 关闭 autoAck
			_ = message.Ack(false)
		}
	}()
	<-forever
}

// ConsumerExchange 接收者 - 带交换机的
func ConsumerExchange(exchange string, types string, routingKey string, queueName string, callback MQCallback) {
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

	//创建exchange
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	// 创建队列 - 临时队列，不需要名字
	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return
	}

	// 绑定这个队列和交换机
	err = channel.QueueBind(queue.Name, routingKey, exchange, false, nil)
	if err != nil {
		return
	}

	// 获取
	messages, err := channel.Consume(queue.Name, "", false,
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
			// 手动应答 - 关闭 autoAck
			_ = message.Ack(false)
		}
	}()
	<-forever
}

// ConsumerDlx 死信队列，常用于定时
func ConsumerDlx(exchangeA string, qAName string, exchangeB string, qBName string, ttl int, callback MQCallback) {
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

	//创建A exchange
	err = channel.ExchangeDeclare(
		exchangeA,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	//创建A queue -
	//1.设置超时时间  x-message-ttl
	//2.设置死信交换机  x-dead-letter-exchange
	//"x-dead-letter-routing-key":"",
	//"x-dead-letter-queue":""
	queueA, err := channel.QueueDeclare(
		qAName,
		true,
		false,
		true,
		false,
		amqp.Table{
			"x-message-ttl":          ttl,
			"x-dead-letter-exchange": exchangeB,
		},
	)
	if err != nil {
		return
	}

	//绑定 exchangeA   queueA
	err = channel.QueueBind(
		queueA.Name, "",
		exchangeA,
		false, nil)
	if err != nil {
		return
	}

	// 重复以上步骤
	//创建B exchange
	err = channel.ExchangeDeclare(
		exchangeB,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	//创建B queue - 这里的队列是一个正常的队列，只是消费A转发过来的消息
	queueB, err := channel.QueueDeclare(
		qBName,
		true,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return
	}

	//绑定 exchangeA   queueA
	err = channel.QueueBind(
		queueB.Name, "",
		exchangeB,
		false, nil)
	if err != nil {
		return
	}

	// 接受消息
	messages, err := channel.Consume(queueB.Name, "",
		false, false,
		false, false, nil)
	if err != nil {
		return
	}

	// 协程接受
	forever := make(chan bool)
	go func() {
		for message := range messages {
			msg := BytesToString(&(message.Body))
			callback(*msg)
			// 回执
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
