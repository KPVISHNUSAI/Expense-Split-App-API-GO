// db/db.go
package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"splitwise-backend/config"
)

var DB *pgxpool.Pool

func ConnectDB(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var err error
	DB, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Check connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = DB.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}
	log.Println("Connected to database!")
}
