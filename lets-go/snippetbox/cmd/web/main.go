package main

import (
	"flag"
	color "github.com/ashidiqidimas/snippetbox/internal"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, color.Blue+"INFO\t"+color.Reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, color.Red+"ERROR\t"+color.Reset, log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	infoLog.Printf("Starting server on %s\n", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
