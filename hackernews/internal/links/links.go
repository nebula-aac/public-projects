package links

import (
	"log"

	database "github.com/nebula-aac/public-projects/hackernews/internal/pkg/db/mysql"
	"github.com/nebula-aac/public-projects/hackernews/internal/users"
)

// #1 Define struct that represents a link
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// #2 Save function that inserts a Link object into database and returns it's ID
func (link Link) Save() int64 {
	//#3 SQL Query to insert link into Links Table
	// Create foreign key from the link we're inserting to the user
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4 Execute SQL statement
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	//#5 Retrieve ID of inserted Link
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID") // changed
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username) // changed
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		} // changed
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
