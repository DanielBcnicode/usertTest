/*
*
FakeMemoryQueue is a basic emulator of rabbitmq only to test purpose
*
*/
package broker

import (
	"errors"

	"github.com/streadway/amqp"
	"usertest.com/event"
)

// FakeMemoryQueue data type to implement an in-memory queue to testing purposes
type FakeMemoryQueue struct {
	Queue []event.DomainEvent
}

// NewFakeMemoryQueue is the constructor of a FakeMemoryQueue
func NewFakeMemoryQueue() *FakeMemoryQueue {
	return &FakeMemoryQueue{Queue: make([]event.DomainEvent, 0)}
}

// Close in this implementation delete all the events from the queue
func (f *FakeMemoryQueue) Close() {
	f.Queue = make([]event.DomainEvent, 0)
}

// Consumer is not implemented
func (r *FakeMemoryQueue) Consumer(queue string) (<-chan amqp.Delivery, error) {
	return nil, nil
}

// PublishDomainEvent put the event in the stack
func (r *FakeMemoryQueue) PublishDomainEvent(event *event.DomainEvent) error {
	r.Queue = append(r.Queue, *event)

	return nil
}

// GetDomainEvent get the event from the stack and delete it
func (r *FakeMemoryQueue) GetDomainEvent() (event.DomainEvent, error) {

	if len(r.Queue) == 0 {
		return event.DomainEvent{}, errors.New("No events in the queue")
	}

	event := r.Queue[len(r.Queue)-1]
	r.Queue = r.Queue[:len(r.Queue)-1]

	return event, nil
}
