# godiator
Godiator is Simple Pub&Sub implementation For Golang.

## Simple Usage Example

```
package main

import "fmt"

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
