package auto_render

import (
    "log"
)

type EventManager struct {
    eventMap map[string]*handlerContaine
    eventChan chan *Event
}

func NewEM() *EventManager {
   em :=  &EventManager{
        eventMap: make(map[string]*handlerContainer),
        eventMap: make(chan *Event, 1)
    }
    go em.process()
}

func (ec *EventManager) addEventListener(name string, h eventHandler) eventHandler {
    if container, ok := ec.eventMap[name]; ok {
        container.addHander(h)
        return h
    }
    container := &handlerContainer{}
    container.AddHandler(h)
    return h
}

func (ec *EventManager) RemoveEventListener(name string, h eventHandler) eventHandler {
    if container, ok := ec.eventMap[name]; ok {
        if h == nil {
            delete(ec.eventMap, name)
        } else {
            container.RemoveHandler(h)
        }
        return h
    }
    return nil
}

func (ec *EventManager) Trigger(name string, data interface{}) *Event{
    event := &Event{name, data})
    ec.eventChan <- event
    return event
}

func (ec *EventManager) process() {
    go (func() {
        for event, ok := <- ec.eventChan {
            if !ok {
                break
            }
            name := event.Name
            if container, ok := ec.eventMap[name]; ok {
                go container.Go(event)
            } else {
                log.Printf("No handler for event %s\n", name)
            }
        }
    })()
}

type Event struct {
    Name string
    Data interface{}
}

type eventHandler func(*Event)

type handlerContainer {
    handlers []eventHandler
}


func (this *handlerContainer) AddHandler(h eventHandler) {
    this.handlers = append(this.handlers, handler)
}

func (this *handlerContainer) RemoveHandler(h eventHandler) {
    for index, handler := range(this.handlers) {
        if handler == h {
            this.handlers = append(this.handlers[:index], this.handlers[index+1:]...)
            break
        }
    }
}

func (this *handlerContainer) Go(e Event) {
    for _, hander := range(this.handlers) {
        handler(e)
    }
}

func init() {
    EM := NewEM()
}