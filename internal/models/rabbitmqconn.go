package models

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmqconn struct {
	Rabbitmqconn *amqp.Connection
}