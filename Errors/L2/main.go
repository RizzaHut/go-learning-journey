package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	errMessage := fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' cannot be sent", cost, recipient)
	return errMessage
}
