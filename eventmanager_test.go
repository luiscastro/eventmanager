package eventmanager

import "testing"

func TestNew(t *testing.T) {
	event := New("test", func(e EventData) EventResponse {
		return EventResponse{"OK" + e.Data.(string)}
	})

	if !Exists("test") {
		t.Error("Event created but not exists?")
	}

	var response EventResponse
	var err error

	response, err = event.Call("1")
	if err != nil {
		t.Error(err)
	}
	if response.Data.(string) != "OK1" {
		t.Error("Event call by object not return correct result")
	}

	response, err = Call("test", "2")
	if err != nil {
		t.Error(err)
	}
	if response.Data.(string) != "OK2" {
		t.Error("Event call not return correct result")
	}

	c := make(chan EventResponse)

	err = event.AsyncCall("3", c)
	if err != nil {
		t.Error(err)
	}
	response = <-c
	if response.Data.(string) != "OK3" {
		t.Error("Event async call by object not return correct result")
	}

	err = AsyncCall("test", "4", c)
	if err != nil {
		t.Error(err)
	}
	response = <-c
	if response.Data.(string) != "OK4" {
		t.Error("Event async call not return correct result")
	}
}
