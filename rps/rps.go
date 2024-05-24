package rps

import (
	"fmt"
	"math/rand"
	"strings"
)

type rpsdata struct {
	Outcome       string `json:"outcome"`
	Winner        string `json:"winner"`
	Loser         string `json:"loser"`
	Winner_choice string `json:"winner_choice"`
	Loser_choice  string `json:"loser_choice"`
}

func Get_User(user_choice string) string {
	user_choice = strings.ToLower(user_choice)
	return user_choice
}

func Validate(user_choice string) bool {
	return user_choice == "rock" || user_choice == "paper" || user_choice == "scissors"
}

func Winner(user_choice string, bot_choice string) rpsdata {
	if user_choice == bot_choice {
		return rpsdata{
			Outcome:       "It's a Tie!",
			Winner:        "None",
			Loser:         "None",
			Winner_choice: user_choice,
			Loser_choice:  bot_choice,
		}
	} else if (user_choice == "rock" && bot_choice == "scissors") || (user_choice == "paper" && bot_choice == "rock") || (user_choice == "scissors" && bot_choice == "paper") {
		return rpsdata{
			Outcome:       "User Wins!",
			Winner:        "User",
			Loser:         "Bot",
			Winner_choice: user_choice,
			Loser_choice:  bot_choice,
		}
	} else {
		return rpsdata{
			Outcome:       "Bot Wins!",
			Winner:        "Bot",
			Loser:         "User",
			Winner_choice: bot_choice,
			Loser_choice:  user_choice,
		}
	}
}

func Play(user_choice string) (rpsdata, error) {
	bot_choice := []string{"rock", "paper", "scissors"}[rand.Intn(3)]
	user_choice = Get_User(user_choice)
	if !Validate(user_choice) {
		return rpsdata{}, fmt.Errorf("not a valid input")
	}
	return Winner(user_choice, bot_choice), nil
}
