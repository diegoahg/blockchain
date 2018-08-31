package main

import (
	api "github.com/diegoahg/blockchain/api"
	"github.com/diegoahg/blockchain/app"
	_ "github.com/diegoahg/blockchain/config"
)

func main() {
	ai := app.New()
	api.Init(ai)
	ai.Run()
}
