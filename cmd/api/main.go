package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-away-learning/ama/internal/database"
	"github.com/go-away-learning/ama/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading environment variables. Err: %v", err)
	}

	ctx := context.Background()
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(fmt.Sprintf("error creating connection pool. Err: %v", err))
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(fmt.Sprintf("error pinging database. Err: %v", err))
	}

	serverAddr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	server := server.New(serverAddr, database.New(pool))
	go func() {
		fmt.Printf("starting server on %s\n", serverAddr)
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(fmt.Sprintf("cannot start server: %v", err))
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
