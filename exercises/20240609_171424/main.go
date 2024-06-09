package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	suggestedFriendsSet := make(map[string]bool)
	usernameDirectFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		usernameDirectFriends[friend] = true
	}

	for directFriend := range usernameDirectFriends {
		for _, friend := range friendships[directFriend] {
			if !usernameDirectFriends[friend] && friend != username {
				suggestedFriendsSet[friend] = true
			}
		}
	}

	var suggestedFriends []string
	for friend := range suggestedFriendsSet {
		suggestedFriends = append(suggestedFriends, friend)
	}

	return suggestedFriends
}
