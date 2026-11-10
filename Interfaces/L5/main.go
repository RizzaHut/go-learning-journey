package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed == false {
		cost := 5 * len(e.body)
		return cost
	} else {
		cost := 2 * len(e.body)
		return cost
	}
}

func (e email) format() string {
	if e.isSubscribed == false {
		new_string := fmt.Sprintf("'%s' | Not Subscribed",e.body )
		return new_string
	} else {
		new_string := fmt.Sprintf("'%s' | Subscribed", e.body)
		return new_string
	}
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
