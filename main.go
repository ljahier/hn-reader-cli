// bubbletea: list mix between paginator and fancy list and when we press enter, open browser with hn page
// upgrade idea, use Tabs to reproduce same schema like official HN website

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const HACKER_NEWS_API_BASE_URL = "https://hacker-news.firebaseio.com/v0"

type HNResponse struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	StoryType   string `json:"type"`
	Url         string `json:"url"`
}

type storyMsg []string

func main() {
	p := tea.NewProgram(model{})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func GetStories() tea.Msg {
	var stories []string
	storiesId := getTopStories()

	for _, storyId := range storiesId[:30] {
		stories = append(stories, fetchStory(storyId).Title)
	}
	return storyMsg(stories)
}

/*
- TODO: handle error in http.Get
*/
func getTopStories() []string {
	resp, _ := http.Get(HACKER_NEWS_API_BASE_URL + "/topstories.json")

	defer resp.Body.Close()
	rawBody, _ := io.ReadAll(resp.Body)
	body := strings.TrimSuffix(strings.TrimPrefix(string(rawBody), "["), "]") // use struct []int
	return strings.Split(body, ",")
}

/*
- TODO: handle error in http.Get
*/
func fetchStory(storyId string) HNResponse {
	resp, _ := http.Get(HACKER_NEWS_API_BASE_URL + "/item/" + storyId + ".json")
	hnResponse := new(HNResponse)

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(hnResponse)
	return *hnResponse
}
