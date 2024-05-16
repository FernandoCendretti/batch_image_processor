package event

// Implementar testes

type Publisher interface {
	Publish(eventName string, payload interface{}) error
}

type Event struct {
	p Publisher
}

func NewEvent(p Publisher) *Event {
	return &Event{p: p}
}

func (e *Event) Producer(message interface{}, event string) error {
	return e.p.Publish(event, message)
}
