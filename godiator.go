// Package godiator is simple Pub&Sub package
// You can simply adds any numbers of events and adds any numbers of handlers
// You can publish events and handles this events with your handlers
package godiator

import (
	"fmt"
)

// NewGodiator Create new godiator instance
func NewGodiator() *Godiator {
	return &Godiator{
		subscriptions: make(map[string][]Handler),
	}
}

// Subscribe events with event name and binds with event handler
// Also this method returns godiator instance you can chain any number of handlers same or different events
func (gb *Godiator) Subscribe(eventName string, eventHandler Handler) *Godiator {
	if currentEventHandlers, ok := gb.subscriptions[eventName]; ok {
		gb.subscriptions[eventName] = append(currentEventHandlers, eventHandler)

		return gb
	}

	gb.subscriptions[eventName] = []Handler{eventHandler}

	return gb
}

// Publish events to correct handlers
// Also this method returns error if it can`t find any handlers for event
func (gb *Godiator) Publish(event Event) error {
	if len(event.Name) <= 0 {
		return fmt.Errorf("Can`t specify event name, maybe you remember adds `Name` parameter")
	}

	if eventHandlers, ok := gb.subscriptions[event.Name]; ok {
		for _, eventHandler := range eventHandlers {
			eventHandler(event)
		}

		return nil
	}

	return fmt.Errorf("Can`t find any event registered as `%s`", event.Name)
}
