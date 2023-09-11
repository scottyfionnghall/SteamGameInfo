package steamgamesinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Type describing game summery from steam
type GameSummary struct {
	Success       int `json:"success"`
	QuerrySummery struct {
		NumReviews      int    `json:"num_reviews"`
		ReviewScore     int    `json:"review_score"`
		ReviewScoreDesc string `json:"review_score_desc"`
		TotalPositive   int    `json:"total_positive"`
		TotalNegative   int    `json:"total_negative"`
		TotalReview     int    `json:"total_reviews"`
	} `json:"query_summary"`
	Reviews []Reviews
	Cursor  string `json:"cursor"`
}

type Author struct {
	SteamId              string `json:"steamid"`
	NumGamesOwned        int    `json:"num_games_owned"`
	NumReviews           int    `json:"num_reviews"`
	PlaytimeForever      int    `json:"playtime_forever"`
	PlayTimeLastTwoWeeks int    `json:"playtime_last_two_weeks"`
	PlaytimeAtReview     int    `json:"playtime_at_review"`
	LastPlayed           int    `json:"last_played"`
}

type Reviews struct {
	RecommendationId         string `json:"recommendationid"`
	Author                   Author
	Language                 string  `json:"language"`
	Review                   string  `json:"review"`
	TimestampCreated         int     `json:"timestamp_created"`
	TimeStampUpdate          int     `json:"timestamp_updated"`
	VotedUp                  bool    `json:"voted_up"`
	VotesUp                  int     `json:"votes_up"`
	VotesFunny               int     `json:"votes_funny"`
	WeightedVoteScore        float64 `json:"weighted_vote_score"`
	CommentCount             int     `json:"comment_count"`
	SteamPurchase            bool    `json:"steam_purchase"`
	RecivedForFree           bool    `json:"recived_for_free"`
	WrittenDuringEarlyAccess bool    `json:"written_during_early_access"`
	HiddenInSteamChina       bool    `json:"hidden_in_steam_china"`
	SteamChinaLocation       string  `json:"steam_china_location"`
}

// Function to create and return GameSummary based on the name of the game
func (g AppList) GetGameSummary(name string) GameSummary {
	var r GameSummary
	index := g.Search(name)
	appid := g.AppList.Apps[index].AppId
	resp, err := http.Get(fmt.Sprint("http://store.steampowered.com/appreviews/", appid, "?json=1"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &r)
	return r
}
