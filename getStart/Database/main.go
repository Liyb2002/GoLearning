package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/go-sql-driver/mysql"
)


//var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   "root",
        Passwd: "",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
	
    // Get a database handle.
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}