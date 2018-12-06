package lsp

import (
	"fmt"
	"testing"
)

func TestPending(t *testing.T) {
	w := Window{
		msgs: make(chan Message),
		get:  make(chan chan []pending),
	}
	w.start()

	ack := NewAck(0, 0)
	w.queue(ack)
	ack = NewAck(1, 1)
	w.queue(ack)
	buff := w.pending()
	pending0, pending1 := buff[0], buff[1]
	before := pending0.time.Before(pending1.time)
	fmt.Println("Buffer : ", buff)
	if !before {
		t.Errorf("messsages not properly queued")
	}

}
