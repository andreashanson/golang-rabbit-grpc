package rabbit

import "github.com/streadway/amqp"

type RabbitRepo struct {
	ch *amqp.Channel
}

func NewRabbit(c *amqp.Connection) *RabbitRepo {
	ch, err := c.Channel()
	if err != nil {
		panic(err)
	}
	return &RabbitRepo{ch: ch}
}

func (r *RabbitRepo) Publish(queue, body string) error {
	return r.ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}

func (r *RabbitRepo) CreateQueue(queue string) error {
	_, err := r.ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)

	return err
}
