package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

func NewRabbitConnection(amqpServerURL string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("%s:5672", amqpServerURL))
}
