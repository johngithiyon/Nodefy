package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateRabbitconn() error {

	rabbitmqconn :=  os.Getenv("RABBITMQ_CONN")

	conn,conerr := amqp.Dial(rabbitmqconn)

	if conerr != nil {
		log.Println("Rabbitmq Connection Failed",conerr)
		return conerr
	}

	defer conn.Close()

	return nil 
}