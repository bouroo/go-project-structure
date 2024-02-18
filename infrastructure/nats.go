package infrastructure

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NatsConn struct {
	ctx       context.Context
	JetStream jetstream.JetStream
	Conn      *nats.Conn
}

type NatsOptions struct {
	nats.Options
}

func (config *NatsOptions) ApplyDefault() *NatsOptions {
	if len(config.Url) == 0 {
		config.Url = nats.DefaultURL
	}
	return config
}

func NewNatsConn(opts NatsOptions) (natsConn *NatsConn, err error) {
	nc, err := nats.Connect(opts.Url)
	if err != nil {
		return
	}
	if nc == nil || !nc.IsConnected() {
		err = nats.ErrInvalidConnection
		return
	}
	return &NatsConn{ctx: context.Background(), Conn: nc}, nil
}

func (n *NatsConn) InitJetStream(opts ...jetstream.JetStreamOpt) (err error) {
	if n.JetStream != nil {
		return
	}
	n.JetStream, err = jetstream.New(n.Conn, opts...)
	return
}
