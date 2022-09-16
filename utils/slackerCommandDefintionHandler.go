package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func SlackerCommandDefintionHandler(botContext slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	year := request.Param("year")
	yob, err := strconv.Atoi(year)
	if err != nil {
		println("ERROR WHILE PARSING")
	}
	age := time.Now().Year() - yob
	r := fmt.Sprintf("Your age is: %d", age)
	response.Reply(r)
}
