package steamgamesinfo

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Type to contain Steam response
type AppList struct {
	AppList struct {
		Apps []Game `json:"apps"`
	} `json:"applist"`
}

// Type to contain basic info about a certian game
type Game struct {
	AppId int    `json:"appid"`
	Name  string `json:"name"`
}

// Function initializes AppList variable and makes a response.json file
// If file exists, read it and return it in a AppList format
// If files does not exists, send HTTP GET request, sort information using Quick Sort algorithm
// and create response.json file and return AppList info
func (response *AppList) GetAppList() {
	if _, err := os.Stat("response.json"); err == nil {
		file, fileopen_err := os.ReadFile("response.json")
		if fileopen_err != nil {
			panic(fileopen_err)
		}
		json.Unmarshal(file, &response)
	} else {
		resp, err := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/v2/")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var data AppList
		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}
		quickSort(data.AppList.Apps, 0, len(data.AppList.Apps)-1)
		response.AppList.Apps = data.AppList.Apps
		file, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		os.WriteFile("response.json", file, 0666)
	}
}

// Binary Search algorithm to find info about entered game
// Returns an index from AppList.Apps to use it later
func (g AppList) Search(x string) int {
	arr := g.AppList.Apps
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := l + (r-l)/2
		if arr[m].Name == x {
			return m
		}
		if arr[m].Name < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
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
