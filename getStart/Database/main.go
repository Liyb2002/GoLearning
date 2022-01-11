package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/go-sql-driver/mysql"
)

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
} 

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

	//Search by name
	albums, _ := albumsByArtist("John Coltrane", db)
	fmt.Printf("Albums found: %v\n", albums)

	//Search by ID
	alb, err := albumByID(2, db)
	if err != nil {
    log.Fatal(err)
		}
	fmt.Printf("Album found: %v\n", alb)

	//Add to database
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
	

}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string, db *sql.DB) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album


    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }

    return albums, nil
}

func albumByID(id int64, db *sql.DB)(Album, error){
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows{
			return alb, fmt.Errorf("no such album")
		}
        return alb, fmt.Errorf("Can't find" )
    }
	return alb, nil

}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album, db *sql.DB) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}