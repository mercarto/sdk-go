package nats

import "github.com/nats-io/nats.go"

// The Subscriber interface allows us to configure how the subscription
type Subscriber interface {
	Subscribe(conn *nats.Conn, subject string) (*nats.Subscription, error)
}

type StandardSubscriber struct {
}

func (s *StandardSubscriber) Subscribe(conn *nats.Conn, subject string) (*nats.Subscription, error) {
	return conn.SubscribeSync(subject)
}

var _ Subscriber = (*StandardSubscriber)(nil)

type QueueSubscriber struct {
	Queue string
}

func (s *QueueSubscriber) Subscribe(conn *nats.Conn, subject string) (*nats.Subscription, error) {
	return conn.QueueSubscribeSync(subject, s.Queue)
}

var _ Subscriber = (*QueueSubscriber)(nil)
