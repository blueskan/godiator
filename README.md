# godiator
Godiator is Simple Pub&Sub implementation For Golang.

### Get Package

```go get github.com/blueskan/godiator```

### Simple Usage Example

```go
package main

import "github.com/blueskan/godiator"

func main() {
	godiator := NewGodiator()

	// Subscription
	godiator.
		Subscribe("notification:user:registration", func(event Event) {
			// E-mail handler

			fmt.Println(event.payload.(string) + ": Sending email...")
		}).
		Subscribe("notification:user:registration", func(event Event) {
			// Sms handler

			fmt.Println(event.payload.(string) + ": Sending sms...")
		})

	// Publish
	godiator.Publish(Event{
		name:    "notification:user:registration",
		payload: "New user created",
	})

}
```
`You can send any type of payload maps, arrays, another structs or even nil (default value)`

### GoDoc ###

https://godoc.org/github.com/blueskan/godiator

### TODO ###
- Async message handling (just for now it's just handling sync)
- Any pull requests or ideas will be appreciated
