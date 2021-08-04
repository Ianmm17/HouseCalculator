package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	username = "root"
	password = "password123"
	hostname = "(127.0.0.1:3306)"
	dbname   = "DebtStorage"
)

func DBSetup(debt string) {
	fmt.Println("GO MySQL Tutorial")
	//DataBase setup
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp%s/%s", username, password, hostname, dbname))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}

	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO DebtStorage.DebtTable (Debt, Date) VALUES (?, ?)")
	if err != nil {
		log.Printf("Error %s when it is preparing", err)
	}
	_, err = stmt.Exec(debt, curdate())
	if err != nil {
		log.Printf("Error %s when inserting DB\n", err)
	}
}

func curdate() interface{} {
	dt := time.Now()
	//Format MM-DD-YYYY
	fmt.Println(dt.Format("01-02-2006"))
	dtFMT := dt.Format("01-02-2006")
	return dtFMT

}
