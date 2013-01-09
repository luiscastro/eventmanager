package eventmanager

import "testing"

func TestNew(t *testing.T) {
	event := New("test", func(d interface{}) (interface{}, error) {
		return "OK" + d.(string), nil
	})

	if !Exists("test") {
		t.Error("Event created but not exists?")
	}

	var response interface{}
	var err error

	response, err = event.Call("1")
	if err != nil {
		t.Error(err)
	}
	if response.(string) != "OK1" {
		t.Error("Event call by object not return correct result")
	}

	response, err = Call("test", "2")
	if err != nil {
		t.Error(err)
	}
	if response.(string) != "OK2" {
		t.Error("Event call not return correct result")
	}

	c := make(chan bool)

	err = event.AsyncCall("3", func(d interface{}, err error) {
		if d.(string) != "OK3" {
			t.Error("Event async call by object not return correct result")
		}
		c <- true
	})
	if err != nil {
		t.Error(err)
	}

	err = AsyncCall("test", "4", func(d interface{}, err error) {
		if d.(string) != "OK4" {
			t.Error("Event async call not return correct result")
		}
		c <- true
	})
	if err != nil {
		t.Error(err)
	}

	<-c
	<-c
}
