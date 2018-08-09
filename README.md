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

			fmt.Println(event.message + ": Sending email...")
		}).
		Subscribe("notification:user:registration", func(event Event) {
			// Sms handler

			fmt.Println(event.message + ": Sending sms...")
		})

	// Publish
	godiator.Publish(Event{
		name:    "notification:user:registration",
		message: "New user created",
	})

}
```

### TODO ###
- Async message handling (just for now it's just handling sync)
