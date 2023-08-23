package natslib

import (
	"github.com/nats-io/nats.go"
)

type client struct {
	nc     *nats.Conn
	server string
}

func NewNatsClient(server string) NatsClient {
	return &client{server: server}
}

func (c *client) Connect() error {
	nc, err := nats.Connect(c.server)
	if err != nil {
		return err
	}
	c.nc = nc
	return nil
}

func (c *client) Publish(subject string, message []byte) error {
	return c.nc.Publish(subject, message)
}

func (c *client) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return c.nc.Subscribe(subject, cb)
}

func (c *client) Close() {
	c.nc.Close()
}
