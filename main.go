package main

import (
	"fmt"
	"steamgamesinfo"
)

func main() {
	fmt.Println("Initializng...")
	var games steamgamesinfo.AppList
	games.GetAppList()
	summary := games.GetGameSummary("VA-11 Hall-A: Cyberpunk Bartender Action")
	fmt.Println(summary.QuerrySummery.ReviewScoreDesc)
}
