package main

import (
	"log"

	"github.com/bloodgroup-cplusplus/Full-Stack-App/internal/env"
)


func main() {

	cfg := config{
		addr: env.GetString("ADDR",":8059"),
	}

	app := &application{
		config: cfg,

	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}