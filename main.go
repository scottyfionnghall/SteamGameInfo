package main

import (
	"fmt"
	"steamgamesinfo/steam"
)

func main() {
	fmt.Println("Initializng...")
	var games steam.AppList
	games.GetAppList()
	summary := games.GetGameSummary("VA-11 Hall-A: Cyberpunk Bartender Action")
	fmt.Println(summary.QuerrySummery.ReviewScoreDesc)
}
