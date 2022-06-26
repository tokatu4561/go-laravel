package main

import (
	"myapp/handlers"
	"myapp/data"

	"github.com/tsawler/celeritas"
)

type application struct {
	App *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}

