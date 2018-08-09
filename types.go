package godiator

// Handler type represent func()Event type represented as EventHandler
type Handler func(Event)

// Event struct is simply contains name as string and payload as any type of value,
// You can send key-value pair, array, simple string or even nil (interface{} default value)
type Event struct {
	Name    string
	Payload interface{}
}

// Godiator struct is representing main struct for Godiator
// Also its store subscriptions
type Godiator struct {
	subscriptions map[string][]Handler
}
