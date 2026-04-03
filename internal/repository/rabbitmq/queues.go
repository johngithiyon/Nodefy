package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func Createqueue(chl *amqp091.Channel) error {

	queues := []string{
		"explore_queue",
		 "build_queue",
		 "deploy_queue",
	}

	for _,queue := range queues { 

	chl.QueueDeclare(
         queue,
		 true,
		 false,
		 false,
		 false,
		 nil,
	)
}	

	return nil 
}