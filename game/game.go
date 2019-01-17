package game

import (
	"fmt"

	"github.com/kalleep/hypertyperbot/client"
)

type Game struct {
	id                 string
	name               string
	client             *client.Client
	multiplier         int
	level              int
	score              int
	increaseMultiplier int
}

func NewGame() *Game {
	return &Game{
		client:     client.NewClient(),
		multiplier: 1,
		name:       "githyperbot",
	}
}

func (g *Game) Start() {

	g.id = g.client.GetGameId()

	highscore := g.client.GetHighscore()

	fmt.Printf("current highscore: %d \n", highscore)

	g.level = 0

	for g.score < highscore+5000 {

		words := g.client.GetWords(g.id, g.level)

		for _, word := range words {

			for range word {
				g.increaseMultiplier += 1

				if g.increaseMultiplier == 10 {
					g.multiplier += 1
					g.increaseMultiplier = 0
				}

			}

			length := len(word)

			g.score += length * g.multiplier

			g.client.SolveWord(g.id, word, g.level, g.multiplier)

			//time.Sleep(time.Duration(2) * time.Second)

		}

		g.level += 1

	}

	g.client.Gameover(g.id)

	g.client.RegisterScore(g.id, g.name, g.score)

	newHighscore := g.client.GetHighscore()

	fmt.Printf("new highscore: %d \n", newHighscore)

}
