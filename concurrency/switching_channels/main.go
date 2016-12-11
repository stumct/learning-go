package main

import "fmt"

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)
	/*
		msg := Message{
			To:      []string{"test@test.com"},
			From:    "person@person.com",
			Content: "This is a test of a channel",
		}

		failedMessage := FailedMessage{
			ErrorMessage:    "Message intercepted",
			OriginalMessage: Message{},
		}

		msgCh <- msg
		errCh <- failedMessage
	*/
	select {
	case revievedMsg := <-msgCh:
		fmt.Println(revievedMsg)
	case revievedErr := <-errCh:
		fmt.Println(revievedErr)
	default:
		fmt.Println("No messages")
	}

}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
