package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func Createchannel() (*amqp091.Channel,error) {

	 chl,chlerr := Rabbitmq.Rabbitmqconn.Channel()

	 if chlerr != nil {
		log.Println("Channel Creation Error",chlerr)
		return nil,chlerr 
	 }

	 return chl,nil 
}