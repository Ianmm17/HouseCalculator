package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	username = "root"
	password = "password123"
	hostname = "(127.0.0.1:3306)"
	dbname   = "DebtStorage"
)

func DBSetup(debt string) {
	fmt.Println("GO MySQL Tutorial")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp%s/%s", username, password, hostname, dbname))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}

	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO DebtStorage.DebtTable (debt) VALUES (?)")
	if err != nil {
		log.Printf("Error %s when it is preparing", err)
	}
	_, err = stmt.Exec(debt)
	if err != nil {
		log.Printf("Error %s when inserting DB\n", err)
	}

}
