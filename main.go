package main

import (
	"flag"
	"fmt"
	"github.com/kalleep/hypertyperbot/game"
)

func main() {

	name := flag.String("name", "asdwe", "name to register score on")
	beatHighscoreWith := flag.Int("score", 1000, "number to beat the highscore with")

	flag.Parse()

	fmt.Println(*name)

	g := game.NewGame()

	g.Start(*name, *beatHighscoreWith)

}
