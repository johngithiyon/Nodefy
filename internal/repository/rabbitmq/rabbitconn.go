package rabbitmq

import (
	"log"
	"os"

	"github.com/johngithiyon/Nodefy/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Rabbitmq models.Rabbitmqconn

func CreateRabbitconn() error {

	rabbitmqconn :=  os.Getenv("RABBITMQ_CONN")

	conn,conerr := amqp.Dial(rabbitmqconn)

	if conerr != nil {
		log.Println("Rabbitmq Connection Failed",conerr)
		return conerr
	}

	defer conn.Close()

	Rabbitmq.Rabbitmqconn = conn

	chl,channelerr := Createchannel()

	if channelerr != nil {
		return channelerr
	}

	queueerr := Createqueue(chl)

	if queueerr != nil {
		return queueerr
	}

	return nil 
}