The Chain of Responsibility is a behavioral design pattern used to pass a request along a chain of handlers. Each handler decides either to process the request or to pass it to the next handler in the chain.

package main

import "fmt"

// Handler defines the interface for handling requests
type Handler interface {
	SetNext(handler Handler)
	Handle(request int)
}

// BaseHandler holds the next handler reference
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) CallNext(request int) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

// Concrete handler for low-level requests
type LowLevelHandler struct {
	BaseHandler
}

func (h *LowLevelHandler) Handle(request int) {
	if request <= 10 {
		fmt.Println("LowLevelHandler handled request:", request)
	} else {
		h.CallNext(request)
	}
}

// Concrete handler for mid-level requests
type MidLevelHandler struct {
	BaseHandler
}

func (h *MidLevelHandler) Handle(request int) {
	if request <= 20 {
		fmt.Println("MidLevelHandler handled request:", request)
	} else {
		h.CallNext(request)
	}
}

// Concrete handler for high-level requests
type HighLevelHandler struct {
	BaseHandler
}

func (h *HighLevelHandler) Handle(request int) {
	fmt.Println("HighLevelHandler handled request:", request)
}

func main() {
	low := &LowLevelHandler{}
	mid := &MidLevelHandler{}
	high := &HighLevelHandler{}

	low.SetNext(mid)
	mid.SetNext(high)

	// Test the chain
	low.Handle(5)   // handled by LowLevelHandler
	low.Handle(15)  // handled by MidLevelHandler
	low.Handle(50)  // handled by HighLevelHandler
}
