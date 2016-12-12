package main

import (
	"fmt"
)

func main() {
	btn := MakeButton()
	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One: " + msg)
		}
	}()
	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two: " + msg)
		}
	}()

	btn.TriggerEvent("click", "ButtonClicked!!!!")

	btn.RemoveEventListener("click", handlerTwo)

	btn.TriggerEvent("click", "Button Clicked Again!!!!")

	fmt.Scanln()

}

// Create a button object which contains a map of channels "eventListeners"
type Button struct {
	eventListeners map[string][]chan string
}

// Constructor function to make a button
func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

func (this *Button) AddEventListener(event string, resposneChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		this.eventListeners[event] = append(this.eventListeners[event], resposneChannel)
	} else {
		this.eventListeners[event] = []chan string{resposneChannel}
	}
}

func (this *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		for idx, _ := range this.eventListeners[event] {
			if this.eventListeners[event][idx] == listenerChannel {
				this.eventListeners[event] = append(this.eventListeners[event][:idx], this.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (this *Button) TriggerEvent(event string, response string) {
	if _, present := this.eventListeners[event]; present {
		for _, handler := range this.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
