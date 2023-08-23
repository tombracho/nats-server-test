package natslib

import "github.com/nats-io/nats.go"

type NatsClient interface {
	Connect() error
	Publish(subject string, message []byte) error
	Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error)
	Close()
}
