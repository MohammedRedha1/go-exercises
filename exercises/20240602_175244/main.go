package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costCustomer, error := sendSMS(msgToCustomer)
	if error != nil {
		return 0, error
	}
	costSpouse, error := sendSMS(msgToSpouse)
	if error != nil {
		return 0, error
	}
	return costCustomer + costSpouse, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
