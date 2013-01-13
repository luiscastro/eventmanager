//Event manager save events related with string identifier and call events by
//identifier sync or async
package eventmanager

import (
	"errors"
	"sync"
)

var events map[string]*Event
var events_mutex *sync.Mutex

type EventCallback func(interface{}) (interface{}, error)
type EventAsyncCallback func(interface{}, error)

//Event object with Identifier and Callback
type Event struct {
	Identifier string
	Callback   EventCallback
}

func init() {
	events = make(map[string]*Event)
	events_mutex = new(sync.Mutex)
}

//Create new event
func New(identifier string, callback EventCallback) *Event {
	events_mutex.Lock()
	defer events_mutex.Unlock()
	event := new(Event)
	event.Identifier = identifier
	event.Callback = callback
	events[identifier] = event
	return events[identifier]
}

func getCallback(identifier string) EventCallback {
	events_mutex.Lock()
	defer events_mutex.Unlock()
	if event, exists := events[identifier]; exists {
		return event.Callback
	}
	return nil
}

//Verifiy if event exists
func Exists(identifier string) bool {
	callback := getCallback(identifier)
	return callback != nil
}

//Call event callback sync based on Event object
func (e *Event) Call(data interface{}) (interface{}, error) {
	return Call(e.Identifier, data)
}

//Call event callback sync
func Call(identifier string, data interface{}) (interface{}, error) {
	callback := getCallback(identifier)
	if callback != nil {
		return callback(data)
	}
	return nil, errors.New("This event identifier isn't registered")
}

//Call event callback async based on Event object
func (e *Event) AsyncCall(data interface{}, f EventAsyncCallback) error {
	return AsyncCall(e.Identifier, data, f)
}

//Call event callback async the response will be send to channel
func AsyncCall(identifier string, data interface{}, f EventAsyncCallback) error {
	callback := getCallback(identifier)
	if callback != nil {
		go func() {
			response, err := callback(data)
			if f != nil {
				f(response, err)
			}
		}()
		return nil
	}
	return errors.New("This event identifier isn't registered")
}
