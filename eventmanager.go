//Event manager save events related with string identifier and call events by
//identifier sync or async
package eventmanager

import (
	"errors"
	"sync"
)

var events map[string]*Event
var events_mutex *sync.Mutex

//Type which will be sended to event function
type EventData struct {
	Data interface{}
}

//Type which will be returned by event function
type EventResponse struct {
	Data interface{}
}

//Event object with Identifier and Callback
type Event struct {
	Identifier string
	Callback   func(EventData) EventResponse
}

func init() {
	events = make(map[string]*Event)
	events_mutex = new(sync.Mutex)
}

//Create new event
func New(identifier string, callback func(EventData) EventResponse) *Event {
	events_mutex.Lock()
	defer events_mutex.Unlock()
	event := new(Event)
	event.Identifier = identifier
	event.Callback = callback
	events[identifier] = event
	return events[identifier]
}

func getCallback(identifier string) func(EventData) EventResponse {
	events_mutex.Lock()
	defer events_mutex.Unlock()
	if event, exists := events[identifier]; exists {
		return event.Callback
	}
	return nil
}

//Call event callback sync based on Event object
func (e *Event) Call(data interface{}) (EventResponse, error) {
	return Call(e.Identifier, data)
}

//Call event callback sync
func Call(identifier string, data interface{}) (EventResponse, error) {
	callback := getCallback(identifier)
	if callback != nil {
		return callback(EventData{Data: data}), nil
	}
	return EventResponse{}, errors.New("This event identifier isn't registered")
}

//Call event callback async based on Event object
func (e *Event) AsyncCall(data interface{}, c chan<- EventResponse) error {
	return AsyncCall(e.Identifier, data, c)
}

//Call event callback async the response will be send to channel
func AsyncCall(identifier string, data interface{}, c chan<- EventResponse) error {
	callback := getCallback(identifier)
	if callback != nil {
		go func() {
			c <- callback(EventData{data})
		}()
		return nil
	}
	return errors.New("This event identifier isn't registered")
}
