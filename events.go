package events

import (
	"sync"
	"sync/atomic"
)

type (
	Map map[string]interface{}

	// Event emitter
	Event struct {
		currID uint64
		sync.Mutex
		listeners []*Listener
	}

	// Listener instance
	Listener struct {
		id     uint64
		fn     func(Map)
		once   bool
		parent *Event
	}
)

// New Event emitter
func New() *Event {
	return &Event{listeners: []*Listener{}}
}

// Drop Event emitter
func (e *Event) Drop() {
	e = nil
}

// Remove listener
func (l *Listener) Remove() {
	l.parent.RemoveListener(l)
}

func (e *Event) addListener(fn func(Map), once bool) (listener *Listener) {
	listener = &Listener{id: atomic.AddUint64(&e.currID, 1), fn: fn, once: once, parent: e}
	e.Lock()
	defer e.Unlock()
	e.listeners = append(e.listeners, listener)
	return
}

// On - create a new listener
func (e *Event) On(fn func(Map)) (listener *Listener) {
	listener = e.addListener(fn, false)
	return
}

// Once - create a new one-time listener
func (e *Event) Once(fn func(Map)) (listener *Listener) {
	listener = e.addListener(fn, true)
	return
}

// RemoveListener - remove event's listener
func (e *Event) RemoveListener(l *Listener) *Event {
	e.Lock()
	defer e.Unlock()
	for i, v := range e.listeners {
		if v.id == l.id {
			// delete without preserving order
			e.listeners[i] = e.listeners[len(e.listeners)-1]
			e.listeners = e.listeners[:len(e.listeners)-1]
			break
		}
	}

	return e
}

// Clear removes all listeners from (all/event)
func (e *Event) Clear() *Event {
	e.Lock()
	defer e.Unlock()
	e.listeners = []*Listener{}
	return e
}

// ListenersCount returns the count of listeners in the speicifed event
func (e *Event) ListenersCount() int {
	e.Lock()
	defer e.Unlock()
	return len(e.listeners)
}

// Emit new event
func (e *Event) Emit(data Map) *Event {
	listeners := []*Listener{}
	e.Lock()
	defer e.Unlock()
	for _, l := range e.listeners {
		dataSend := make(Map, len(data))
		for k := range data {
			dataSend[k] = data[k]
		}
		l.fn(dataSend)
		if !l.once {
			listeners = append(listeners, l)
		}
	}
	e.listeners = listeners
	return e
}
