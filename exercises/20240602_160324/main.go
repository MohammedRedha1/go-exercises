package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return 2 * len(e.body)
	}
	return 5 * len(e.body)
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | %s", e.body, "Subscribed")
	}
	return fmt.Sprintf("'%s' | %s", e.body, "Not Subscribed")
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
