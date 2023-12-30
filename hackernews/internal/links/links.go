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
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4 Execute SQL statement
	res, err := stmt.Exec(link.Title, link.Address)
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
	stmt, err := database.Db.Prepare("select id, title, address from Links")
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
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
