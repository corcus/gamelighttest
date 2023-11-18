package shared

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RmqClient struct {
	Uri     string
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (client *RmqClient) Connect() error {
	conn, err := amqp.Dial(client.Uri)
	if err != nil {
		return err
	}

	client.conn = conn
	ch, err := client.conn.Channel()
	if err != nil {
		return err
	}
	client.channel = ch
	return nil
}

func (client *RmqClient) Publish(message PublishableMessage) error {

	err := client.channel.Publish(message.Exchange, message.Key, message.Mandatory, message.Immediate, message.Publishing)
	if err != nil {
		return fmt.Errorf("publish message error %w", err)
	}
	return nil
}

type PublishableMessage struct {
	Exchange   string
	Key        string
	Mandatory  bool
	Immediate  bool
	Publishing amqp.Publishing
}

// Consume returns a go channel to receive messages from the requested queue or an error if that fails.
func (client *RmqClient) Consume(queue string) (<-chan amqp.Delivery, error) {
	return client.channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil)
}
