package steam

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Function initializes AppList variable and makes a response.json file
// If file exists, read it and return it in a AppList format
// If files does not exists, send HTTP GET request, sort information using Quick Sort algorithm
// and create response.json file and return AppList info
func GetAppList() (*AppList, error) {
	var response AppList
	if _, err := os.Stat("response.json"); err == nil {
		file, fileopen_err := os.ReadFile("response.json")
		if fileopen_err != nil {
			return nil, fileopen_err
		}
		json.Unmarshal(file, &response)
	} else {
		resp, err := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/v2/")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var data AppList
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
		quickSort(data.AppList.Apps, 0, len(data.AppList.Apps)-1)
		response.AppList.Apps = data.AppList.Apps
		file, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		os.WriteFile("response.json", file, 0666)
	}
	return &response, nil
}

// Binary Search algorithm to find appid of a game.
// Returns int if game found, else returns 0
func (g AppList) GetSteamAppId(x string) (int, error) {
	arr := g.AppList.Apps
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := l + (r-l)/2
		if arr[m].Name == x {
			return g.AppList.Apps[m].AppId, nil
		}
		if arr[m].Name < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return 0, errors.New("Couldn't find this game")
}

// Function to create and return GameSummary based on the name of the game
func (g AppList) GetGameSummary(name string) (*GameSummary, error) {
	var r GameSummary
	appid, err := g.GetSteamAppId(name)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(fmt.Sprint("http://store.steampowered.com/appreviews/", appid, "?json=1"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &r)
	return &r, nil
}

// Part of a QuickSort algorithm
func swap(a *Game, b *Game) {
	t := *a
	*a = *b
	*b = t
}

// Part of a QuickSort algorithm
func partition(array []Game, low int, high int) int {
	pivot := array[high].Name
	i := low - 1
	for j := low; j <= high-1; j++ {
		if array[j].Name < pivot {
			i++
			swap(&array[i], &array[j])
		}
	}
	swap(&array[i+1], &array[high])
	return i + 1
}

// QuickSort algorithm to sort AppList to use Binary Search later
func quickSort(array []Game, low int, high int) {
	if low < high {
		pi := partition(array, low, high)
		quickSort(array, low, pi-1)
		quickSort(array, pi+1, high)
	}
}
