package utils

import (
	"fmt"

	"github.com/shomali11/slacker"
)

func PrintCommandEvents(channel <-chan *slacker.CommandEvent) {
	for event := range channel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
