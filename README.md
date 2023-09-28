# Steam Info in Golang

Packages that can be used to get information about games in Steam.
Usess json files, http requests, quick sort algorithm and binary search algorithm.


## Types

### type Game
```go
type Game struct {
	AppId int    `json:"appid"`
	Name  string `json:"name"`
}
```


### type AppList

```go
type AppList struct {
	AppList struct {
		Apps []Game `json:"apps"`
	} `json:"applist"`
}
```

### func(AppList) GetAppList

```go
func (a *AppList) GetAppList()
```

Initialize variable of type AppList and make response.json file. Function checks if response.json file exists, if not we send request to SteamWebAPI to get list of games and their ID. We create such file so we can later use it than sending each time new request.

### func(AppList) Search
```go
func (a AppList) Search(x string) int
```

Using Binary Search algorithm returns ID of entered game title.

### func(AppList) GetGameSummary

```go
func (g AppList) GetGameSummary(name string) GameSummary 
```

Returns a variable of type GameSummary.

### type GameSummary

```go
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
```

### type Author

```go
type Author struct {
	SteamId              string `json:"steamid"`
	NumGamesOwned        int    `json:"num_games_owned"`
	NumReviews           int    `json:"num_reviews"`
	PlaytimeForever      int    `json:"playtime_forever"`
	PlayTimeLastTwoWeeks int    `json:"playtime_last_two_weeks"`
	PlaytimeAtReview     int    `json:"playtime_at_review"`
	LastPlayed           int    `json:"last_played"`
}
```

type Reviews

```go
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
```


