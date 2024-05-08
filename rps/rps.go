package rps

import (
	"math/rand"
	"strings"
)

func Get_User(user_choice string) string {
	user_choice = strings.ToLower(user_choice[0:1])
	return user_choice
}

func Validate(user_choice string) bool {
	return user_choice == "r" || user_choice == "p" || user_choice == "s"
}

func Winner(user_choice string, bot_choice string) string {
	if user_choice == bot_choice {
		return "It's a draw!"
	} else if (user_choice == "r" && bot_choice == "s") || (user_choice == "p" && bot_choice == "r") || (user_choice == "s" && bot_choice == "p") {
		return "User wins!"
	} else {
		return "Bot wins!"
	}
}

func Play(user_choice string) string {
	bot_choice := []string{"r", "p", "s"}[rand.Intn(3)]
	user_choice = Get_User(user_choice)
	if !Validate(user_choice) {
		return "not a real choice"
	}
	return Winner(user_choice, bot_choice)
}
