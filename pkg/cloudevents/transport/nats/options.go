package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

// Option is the function signature required to be considered an nats.Option.
type Option func(*Transport) error

// applyDefaultOptions ensures that any required fields on the Transport are set
func applyDefaultOptions(t *Transport) error {
	t.Subscriber = &StandardSubscriber{}
	return nil
}

// WithEncoding sets the encoding for NATS transport.
func WithEncoding(encoding Encoding) Option {
	return func(t *Transport) error {
		t.Encoding = encoding
		return nil
	}
}

// WithConn allows callers to set the connection rather than the transport creating it's own connection
func WithConn(conn *nats.Conn) Option {
	return func(t *Transport) error {
		if conn == nil {
			return fmt.Errorf("WithConn(conn): conn must not be nil")
		}

		if conn.IsClosed() {
			return fmt.Errorf("WithConn(conn): conn must not be closed")
		}
		t.Conn = conn

		// a caller may invoke nats.New() with a URL different to the provided connection,
		// we treat the option as overriding behaviour so replace the provided URL
		t.NatsURL = conn.ConnectedUrl()
		return nil
	}
}

// WithConnOptions supplies NATS connection options that will be used when setting
// up the internal NATS connection
func WithConnOptions(opts ...nats.Option) Option {
	return func(t *Transport) error {
		for _, o := range opts {
			t.ConnOptions = append(t.ConnOptions, o)
		}

		return nil
	}
}

// WithQueueSubscriber configures the transport to create a Queue subscription instead of a standard subscription
func WithQueueSubscriber(queue string) Option {
	return func(t *Transport) error {
		t.Subscriber = &QueueSubscriber{queue}
		return nil
	}
}
