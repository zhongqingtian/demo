package rabbitmq

import "github.com/Braveheart7854/rabbitmqPool"
func P()  {
	rabbitmqPool.AmqpServer = rabbitmqPool.Service{

		AmqpUrl:"amqp://guest:guest@localhost:5672/",

		ConnectionNum:10,

		ChannelNum:100,

	}

	rabbitmqPool.InitAmqp()

	message,err := rabbitmqPool.AmqpServer.PutIntoQueue(ExchangeName,RouteKey,data)

	if err != nil{

		//若有错误，则表示消息发送失败，做好失败消息处理

		rabbitmqPool.Logger.Notice(message)

	}

}