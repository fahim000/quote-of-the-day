package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var quotes = []string{
		"Don't die.",
		"Latinas are hot.",
		"I support lesbians.",
		"Whiskey = who is key.",
		"Live, laugh, love.",
		"Go touch some grass.",
		"Sleep is for the weak.",
		"Code. Eat. Repeat.",
		"404: Motivation not found.",
		"Trust the process.",
		"Reality is a suggestion.",
		"Your vibe attracts your tribe.",
		"Stay curious.",
		"Make it work, make it right, make it fast.",
		"Less talk, more rock.",
		"Messy code is better than no code.",
		"Break things. Fix them later.",
		"Commit, push, pray.",
		"Eat bugs. Debug code.",
		"Go lang, go long!",
	}

	fmt.Println(quotes[rand.Intn(len(quotes))])

}
