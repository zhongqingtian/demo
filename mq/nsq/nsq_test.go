package nsq

import (
	"testing"
)

func TestSender(t *testing.T) {
	Sender("test")
}

func TestReceiver1(t *testing.T) {
	Receiver("test", "ch1")
}

func TestReceiver2(t *testing.T) {
	Receiver("test", "ch1")
}

func TestReceiver3(t *testing.T) {
	Receiver("test", "ch2")
}