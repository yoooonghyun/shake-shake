package event

type EventRecorder struct {
	events []interface{}
}

func (er *EventRecorder) AddEvent(event interface{}) { er.events = append(er.events, event) }

func (er *EventRecorder) Clear() {
	er.events = []interface{}{}
}

func (er *EventRecorder) GetEvents() []interface{} { return er.events }
