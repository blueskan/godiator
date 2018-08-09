package godiator

// Handler type represent func()Event type represented as EventHandler.
type Handler func(Event)

// Event struct is simply contains name as string and payload as any type of value,
// you can send maps, arrays, simple strings or even nil (interface{} default value)
type Event struct {
	Name    string
	Payload interface{}
}

// Godiator struct is representing main struct for Godiator,
// also its store subscriptions.
type Godiator struct {
	subscriptions map[string][]Handler
}
