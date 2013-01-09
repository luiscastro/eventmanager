Event Manager for Go
========

This is a simple package written in golang to manage your events based on string identifier.

How it works
--------------

Run ``go get github.com/luiscastro/eventmanager`` to download package

It's realy simple to use just follow the next example:

````
func main() {
  eventmanager.New("github_test", func(d interface{}) (interface{}, error) {
    fmt.Println("My github_test event is working :)")
    return "My Response", nil
  })

  response, err := eventmanager.Call("github_test", nil)
  if err == nil {
    fmt.Println("My event response is:", response)
  }

  err = eventmanager.AsyncCall("github_test", nil, func(d interface{}, err error) {
    fmt.Println("My event response is:", d)
  })
}
```

