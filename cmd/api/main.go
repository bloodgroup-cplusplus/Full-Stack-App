package main

import (
	"log"

	"github.com/bloodgroup-cplusplus/Full-Stack-App/internal/env"
	"github.com/bloodgroup-cplusplus/Full-Stack-App/internal/store"
)


func main() {

	cfg := config{
		addr: env.GetString("ADDR",":8059"),
		db: dbConfig{
			addr : env.GetString("DB_ADDR","postgres://user:adminpassword@loaclhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS",30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS",30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME","15min"),
		},
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