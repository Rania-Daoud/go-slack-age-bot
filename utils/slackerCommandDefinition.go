package utils

import "github.com/shomali11/slacker"

var SlackerCommandDefintion = &slacker.CommandDefinition{
	Description: "YOB Calculator",
	Examples:    []string{"My yob is: 1999"},
	Handler:     SlackerCommandDefintionHandler,
}
