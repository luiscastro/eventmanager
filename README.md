Event Manager for Go
========

This is a simple package written in golang to manage your events based on string identifier.

How it works
--------------

Run ``go get github.com/luiscastro/eventmanager`` to download package

It's realy simple to use just follow the next example:

````
func main() {
  eventmanager.New("github_test", func(e eventmanager.EventData) eventmanager.EventResponse {
    fmt.Println("My github_test event is working :)")
    return eventmanager.EventResponse{"My Response"}
  })

  response, err := eventmanager.Call("github_test", nil)
  if err == nil {
    fmt.Println("My event response is:", response.Data)
  }

  rchan := make(chan eventmanager.EventResponse)
  err = eventmanager.AsyncCall("github_test", nil, rchan)
  if err == nil {
    response = <-rchan
    fmt.Println("My event response is:", response.Data)
  }
}
```

