package main

import (
	"log"
	"social/internal/store"
)

func main() {
	cfg := config{
		addr: ":8085",
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
