package main

import (
	"fmt"

	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/server"
	"github.com/IainMcl/HereWeGo/internal/settings"
	"github.com/IainMcl/HereWeGo/internal/util"
)

func init() {
	settings.Setup()
	logging.Setup()
	util.Setup()
}

func main() {

	logging.Info("Starting server...")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
