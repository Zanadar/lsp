package lsp

import "testing"

func TestPending(t *testing.T) {
	pending := Pending{}

	ack := NewAck(0, 0)
	pending.queue(ack)

}
