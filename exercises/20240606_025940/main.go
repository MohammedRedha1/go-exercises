package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	_, isUser := users[name]
	if !isUser {
		return false, errors.New("not found")
	}
	if users[name].scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
