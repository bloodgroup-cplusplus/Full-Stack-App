package main

import (
	"log"

	"github.com/bloodgroup-cplusplus/Full-Stack-App/internal/env"
	"github.com/bloodgroup-cplusplus/Full-Stack-App/internal/store"
)


func main() {

	cfg := config{
		addr: env.GetString("ADDR",":8059"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:store,


	}

	
	log.Fatal(store)

	mux := app.mount()
	log.Fatal(app.run(mux))

}