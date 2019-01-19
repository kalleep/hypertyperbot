package game

import (
	"log"
	"time"

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
	}
}

func (g *Game) Start(name string, beatHighscoreWith int) {

	g.name = name

	for {
		highscore := g.client.GetHighscore()

		log.Printf("current highscore: %d \n", highscore.Score)
		log.Printf("held by: %s", highscore.Name)

		if highscore.Name != g.name {
			g.id = g.client.GetGameId()

			g.level = 0

			for g.score < highscore.Score+beatHighscoreWith {

				words := g.client.GetWords(g.id, g.level)

				for _, word := range words {

					for range word {
						g.increaseMultiplier += 1

						if g.increaseMultiplier == 10 && g.multiplier < 10 {
							g.multiplier += 1
							g.increaseMultiplier = 0
						}

					}

					length := len(word)

					g.score += length * g.multiplier

					g.client.SolveWord(g.id, word, g.level, g.multiplier)

				}

				g.level += 1

			}

			g.client.Gameover(g.id)

			g.client.RegisterScore(g.id, g.name, g.score)

			newHighscore := g.client.GetHighscore()

			log.Printf("new highscore: %d \n", newHighscore.Score)
			log.Printf("held by: %s \n", newHighscore.Name)
		}

		time.Sleep(time.Duration(10) * time.Minute)
	}

}
