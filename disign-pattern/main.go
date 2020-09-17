package main

import (
	"fmt"
	"unicode"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(100000000000), profile.ProfilePath(".")).Stop()

	h := NewSymbolHandler(
		NewCharacterHandler(
			NewDigitHandler(nil)))

	fmt.Println("---- 1234 ----")
	h.HandlerAction("1234")
	fmt.Println("---- victor ----")
	h.HandlerAction("victor")
}

type Action interface {
	HandlerAction(c string)
}

type handler struct {
	next Action
	Action
}

func (h *handler) SetNext(next Action) {
	h.next = next
}

func (h *handler) toNext(c string) {
	// next part
	if h.next != nil {
		h.next.HandlerAction(c)
	} else {
		fmt.Println("the last handler")
	}
}

type SymbolHandler struct{ handler }

func NewSymbolHandler(next Action) *SymbolHandler {
	return &SymbolHandler{handler: handler{next: next}}
}

func (sh *SymbolHandler) HandlerAction(c string) {
	fmt.Println("Is Symbol")
	sh.toNext(c)
}

type CharacterHandler struct{ handler }

func NewCharacterHandler(next Action) *CharacterHandler {
	return &CharacterHandler{handler: handler{next: next}}
}

func onlyLetter(c string) bool {
	for _, r := range c {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (ch *CharacterHandler) HandlerAction(c string) {
	if onlyLetter(c) {
		fmt.Println("Is Character")
	}
	ch.toNext(c)
}

type DigitHandler struct{ handler }

func NewDigitHandler(next Action) *DigitHandler {
	return &DigitHandler{handler: handler{next: next}}
}

func onlyDigit(c string) bool {
	for _, r := range c {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func (dh *DigitHandler) HandlerAction(c string) {
	if onlyDigit(c) {
		fmt.Println("Is Number")
	}
	dh.toNext(c)
}
