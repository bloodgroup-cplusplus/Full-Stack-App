package main

import (
	"fmt"

	"github.com/bubble-bot/weavebox"
)

func main() {
	app := weavebox.New()
	app.Get("/product",func(*weavebox.Context) error {return nil})
	app.Serve(3001)
	fmt.Println("It's all good")
	fmt.Println("This line is added")
}
