package main

import (
	"log"
	"social/internal/env"
	"social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8085"),
	}
	s := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store:  s,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
