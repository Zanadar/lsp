// Contains the implementation of a LSP client.

package lsp

import (
	"errors"
	"time"

	"github.com/cmu440/lspnet"
)

type client struct {
	// TODO: implement this!
	// pending list : sent and no ack received
	params *Params
}

// NewClient creates, initiates, and returns a new client. This function
// should return after a connection with the server has been established
// (i.e., the client has received an Ack message from the server in response
// to its connection request), and should return a non-nil error if a
// connection could not be made (i.e., if after K epochs, the client still
// hasn't received an Ack message from the server in response to its K
// connection requests).
//
// hostport is a colon-separated string identifying the server's host address
// and port number (i.e., "localhost:9999").
func NewClient(hostport string, params *Params) (Client, error) {
	addr, err := lspnet.ResolveUDPAddr("udp", hostport)
	if err != nil {
		return &client{}, err
	}
	conn, err := lspnet.DialUDP("udp", nil, addr)
	if err != nil {
		return &client{}, err
	}
	msgs := make(chan *Message)
	go func() {
		select {
		case msg := <-msgs:
			msg.String()
			conn.Write([]byte{})
		}
	}()

	// create a connect message and send
	// wait for ack
	// store id from ack
	return nil, errors.New("not yet implemented")
}

func (c *client) ConnID() int {
	return -1
}

func (c *client) Read() ([]byte, error) {
	// TODO: remove this line when you are ready to begin implementing this method.
	select {} // Blocks indefinitely.
	return nil, errors.New("not yet implemented")
}

func (c *client) Write(payload []byte) error {
	return errors.New("not yet implemented")
}

func (c *client) Close() error {
	return errors.New("not yet implemented")
}

type pending struct {
	time time.Time
	msg  Message
}

type Window struct {
	inflight int
	wFront   int
	wBack    int
	buff     []pending
	msgs     chan Message
	get      chan chan []pending
	done     chan struct{}
}

func (w *Window) start() {
	go func() {
		for {
			select {
			case msg := <-w.msgs:
				w.buff = append(w.buff, pending{time.Now(), msg})
			case get := <-w.get:
				get <- w.buff
			}
		}
	}()
}

func (w *Window) queue(m *Message) {
	w.msgs <- *m
}

func (w *Window) pending() []pending {
	get := make(chan []pending)
	w.get <- get
	return <-get
}
