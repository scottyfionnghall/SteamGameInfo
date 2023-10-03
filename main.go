package main

// Example of how to use this.

import (
	"fmt"
	"steamgamesinfo/steam"
)

func main() {
	// Initialize AppList variable.
	fmt.Println("Initializng...")
	games, err := steam.GetAppList()
	if err != nil {
		fmt.Println(err)
		return
	}
	// usign GetGameSummary method, we pass a string with game name.
	summary, err := games.GetGameSummary("Dota 2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(summary.QuerrySummery.ReviewScore)
}
