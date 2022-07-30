package main

import (
	"facade/services/notification"
	"log"
)

func main() {

	// Send function a Facade function that handles all complexity of notifying clients
	err := notification.Send("notification message", "recipient")
	if err != nil {
		log.Println(err)
	}
}
