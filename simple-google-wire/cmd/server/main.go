package main

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nebula-aac/public-projects/simple-google-wire/internal/user"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userHandler := user.Wire(db)
	http.Handle("/user", userHandler.FindByUsername())
	http.ListenAndServe(":8000", nil)
}
