package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kalleep/hypertyperbot/util"
)

type Client struct {
	baseUrl    string
	httpClient *http.Client
}

func NewClient() *Client {

	return &Client{
		baseUrl: "https://server.hypertyper.io",
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) GetGameId() string {

	body := strings.NewReader("")

	request, err := http.NewRequest("POST", c.baseUrl+"/game", body)

	util.PanicIfErr(err)

	response, err := c.httpClient.Do(request)

	util.PanicIfErr(err)

	decoder := json.NewDecoder(response.Body)

	var id string

	err = decoder.Decode(&id)

	util.PanicIfErr(err)

	return id

}

func (c *Client) GetWords(id string, round int) Words {

	url := c.baseUrl + "/words/" + id + "/" + strconv.Itoa(round)

	request, err := http.NewRequest("GET", url, strings.NewReader(""))

	util.PanicIfErr(err)

	response, err := c.httpClient.Do(request)

	util.PanicIfErr(err)

	decoder := json.NewDecoder(response.Body)

	var words Words

	err = decoder.Decode(&words)

	util.PanicIfErr(err)

	return words

}

type scoreRequest struct {
	Word         string `json:"word"`
	CurrentLevel int    `json:"currentLevel"`
	Multiplier   int    `json:"multiplier"`
}

func (c *Client) SolveWord(id, word string, currentLevel, multiplier int) {

	url := c.baseUrl + "/game/" + id + "/score"

	body, err := json.Marshal(scoreRequest{Word: word, CurrentLevel: currentLevel, Multiplier: multiplier})

	util.PanicIfErr(err)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))

	request.Header.Set("Content-Type", "application/json")

	util.PanicIfErr(err)

	_, err = c.httpClient.Do(request)

	util.PanicIfErr(err)
}

func (c *Client) Gameover(id string) {

	url := c.baseUrl + "/game/" + id + "/gameover"

	request, err := http.NewRequest("OPTIONS", url, strings.NewReader(""))

	util.PanicIfErr(err)

	c.httpClient.Do(request)

	request, err = http.NewRequest("PUT", url, strings.NewReader(""))

	util.PanicIfErr(err)

	c.httpClient.Do(request)

}

type finalScore struct {
	ScoreData scoreData `json:"scoreData"`
}

type scoreData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (c *Client) RegisterScore(id, name string, score int) {

	url := c.baseUrl + "/highscore"

	final := finalScore{
		ScoreData: scoreData{
			Id:    id,
			Name:  name,
			Score: score,
		},
	}

	body, err := json.Marshal(final)

	util.PanicIfErr(err)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))

	request.Header.Set("Content-Type", "application/json")

	util.PanicIfErr(err)

	_, err = c.httpClient.Do(request)

	util.PanicIfErr(err)
}

type highscores struct {
	Score []highscore `json:"score"`
}

type highscore struct {
	Id    string `json:"_id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (c *Client) GetHighscore() highscore {

	url := c.baseUrl + "/highscore"

	response, err := c.httpClient.Get(url)

	util.PanicIfErr(err)

	decoder := json.NewDecoder(response.Body)

	var h highscores

	decoder.Decode(&h)

	return h.Score[0]

}

type Words []string
