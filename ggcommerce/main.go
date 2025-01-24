package main

import (
	"fmt"

	"github.com/anthdm/weavebox"
)

func main() {
	app := weavebox.New()
	app.Get("/product",func(*weavebox.Context) error {return nil})
	app.Serve(3001)
	fmt.Println("It's all good")
}