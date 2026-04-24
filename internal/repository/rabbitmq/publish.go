package rabbitmq

import "github.com/rabbitmq/amqp091-go"


func Publish(chl *amqp091.Channel,queuename string,body []byte) error {
    
	
   publisherr := chl.Publish("",queuename,false,false,amqp091.Publishing{
	  ContentType: "application/json",
	  Body:body,
   })

   if publisherr != nil {
	  return publisherr
   }


	return nil 
}