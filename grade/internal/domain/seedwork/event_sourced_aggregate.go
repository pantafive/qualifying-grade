package seedwork

import (
	"fmt"
)

type PersistentDomainEvent interface {
	DomainEvent
	StreamId() fmt.Stringer
	SetStreamId(value fmt.Stringer)
	StreamPosition() uint
	SetStreamPosition(uint)
	StreamType() string
	EventType() string
	EventVersion() uint8
}

type PersistentDomainEventHandler func(event PersistentDomainEvent)
type StreamIdAccessor func() fmt.Stringer

type EventSourcedAggregate struct {
	streamId fmt.Stringer
	handlers map[string]PersistentDomainEventHandler
	EventiveEntity
	VersionedAggregate
}

func (a EventSourcedAggregate) StreamId() fmt.Stringer {
	return a.streamId
}

func (a *EventSourcedAggregate) SetStreamId(val fmt.Stringer) {
	a.streamId = val
}

func (a *EventSourcedAggregate) AddHandler(e PersistentDomainEvent, handler PersistentDomainEventHandler) {
	a.handlers[e.EventType()] = handler
}

func (a *EventSourcedAggregate) LoadFrom(pastEvents []PersistentDomainEvent) {
	for i := range pastEvents {
		a.handlers[pastEvents[i].EventType()](pastEvents[i])
		a.SetVersion(pastEvents[i].StreamPosition())
	}
}

func (a *EventSourcedAggregate) Update(e PersistentDomainEvent) {
	e.SetStreamId(a.StreamId())
	e.SetStreamPosition(a.Version() + 1)
	a.handlers[e.EventType()](e)
	a.SetVersion(a.Version() + 1)
	a.AddDomainEvent(e)
}
