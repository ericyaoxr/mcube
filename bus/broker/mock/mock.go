package broker

import (
	"fmt"

	"github.com/ericyaoxr/mcube/bus"
	"github.com/ericyaoxr/mcube/bus/event"
)

// NewBroker todo
func NewBroker() bus.Publisher {
	return &mockBroker{}
}

type mockBroker struct {
}

func (m *mockBroker) Pub(topic string, e *event.Event) error {
	fmt.Println(topic, e)
	return nil
}
