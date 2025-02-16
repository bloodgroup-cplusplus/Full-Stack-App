package main

import "log"

func main() {

	cfg := config{
		addr: ":8069",
	}

	app := &application{
		config: cfg,
	}
	log.Fatal(app.run())

}