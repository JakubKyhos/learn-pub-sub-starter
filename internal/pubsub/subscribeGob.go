package pubsub

import (
	"bytes"
	"encoding/gob"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeGob[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
	handler func(T) Acktype,
) error {
	return subscribe(
		conn,
		exchange,
		queueName,
		key,
		queueType,
		handler,
		func(data []byte) (T, error) {
			buffer := bytes.NewBuffer(data)
			decoder := gob.NewDecoder(buffer)
			var target T
			err := decoder.Decode(&target)
			return target, err
		},
	)
}
