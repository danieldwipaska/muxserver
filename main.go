package main

import (
	"github.com/danieldwipaska/muxserver/src/routes"
	"github.com/danieldwipaska/muxserver/src/utils"
)

var Movies []utils.Movie

func main() {

	routes.SetupRoutes()

}
