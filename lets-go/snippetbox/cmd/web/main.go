package main

import (
	"database/sql"
	"flag"
	color "github.com/ashidiqidimas/snippetbox/internal"
	"github.com/ashidiqidimas/snippetbox/internal/models"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	//dsn := flag.String("dsn", "root:Ngolo7golo!!@/snippetbox?parseTime=true", "MySQL data source name")

	infoLog := log.New(os.Stdout, color.Blue+"INFO\t"+color.Reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, color.Red+"ERROR\t"+color.Reset, log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer func() {
		err = db.Close()
		errorLog.Println(err)
	}()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	infoLog.Printf("Starting server on %s\n", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
