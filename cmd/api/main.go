package main

import (
	"log"
	"social/internal/db"
	"social/internal/env"
	"social/internal/store"
)

const version = "0.0.1"

// @title Gopher Social API
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in				header
// @name				Authorization
// @description
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8085"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgresql://postgres:admin@localhost:5432/postgres?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	database, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to database")
	s := store.NewStorage(database)
	app := &application{
		config: cfg,
		store:  s,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
