package patterns

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("Handled by handler A")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	} else {
		fmt.Println("Unable to handle the request")
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("Handled by handler B")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	} else {
		fmt.Println("Unable to handle the request")
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerA.SetNext(handlerB)
	handlerA.HandleRequest("A")
	handlerA.HandleRequest("B")
	handlerA.HandleRequest("C")
}
