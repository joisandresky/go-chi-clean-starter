package main

import (
	"github.com/joisandresky/go-chi-clean-starter/internal/infra"
)

func main() {
	server := infra.BuildServer()

	server.Run()
}
