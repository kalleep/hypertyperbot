package main

import (
	"flag"

	"github.com/kalleep/hypertyperbot/game"
)

func main() {

	name := flag.String("name", "hypertyperbot", "name to register score on")
	beatHighscoreWith := flag.Int("score", 1000, "number to beat the highscore with")

	g := game.NewGame()

	g.Start(*name, *beatHighscoreWith)

}
