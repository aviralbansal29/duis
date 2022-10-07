package main

import (
	"github.com/aviralbansal29/duis/config"
	_ "github.com/aviralbansal29/duis/docs/swagger"
	server "github.com/aviralbansal29/duis/transport"
)

// @title Dynamic UI Server
// @version 1.0
// @BasePath /
func main() {
	config.InitiateGlobalInstance()

	s := server.RestServer{}
	s.Setup()
	s.Start()
}
