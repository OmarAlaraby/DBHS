package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"DBHS/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// ---- DB connection ---- //

	dbServerURL := fmt.Sprintf("postgres://postgres:%s@%s:%d/%s?sslmode=disable", config.Password, config.Host, config.Port, config.User)

	dbPool, err := pgxpool.New(context.Background(), dbServerURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	if err = dbPool.Ping(context.Background()); err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	// ---- http server ---- //

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	config.Init(infoLog, errorLog)

	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	defineURLs()

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  config.Mux,
	}

	infoLog.Printf("starting server on :%s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
