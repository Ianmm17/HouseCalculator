package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type UserDebtInfo struct {
	Debt string
	Date string
}

const (
	username = "root"
	password = "password123"
	hostname = "(127.0.0.1:3306)"
	dbname   = "DebtStorage"
)

func DBSetup() (db *sql.DB) {
	//DataBase setup
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp%s/%s", username, password, hostname, dbname))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}
	return db
}
func DBInsert(debt string) {
	//DataBase Inserting Data
	db := DBSetup()
	stmt, err := db.Prepare("INSERT INTO DebtStorage.DebtTable (Debt, Date) VALUES (?, ?)")
	if err != nil {
		log.Printf("Error %s when it is preparing", err)
	}
	_, err = stmt.Exec(debt, curdate())
	if err != nil {
		log.Printf("Error %s when inserting DB\n", err)
	}
}

func DbSelectQuery() []UserDebtInfo {
	db := DBSetup()
	rows, err := db.Query(`SELECT * FROM DebtStorage.DebtTable`)
	if err != nil {
		log.Println(err)
	}
	debtTable := UserDebtInfo{}
	debtsInfo := []UserDebtInfo{}

	for rows.Next() {
		var debt, date string
		err = rows.Scan(&debt, &date)
		if err != nil {
			log.Println(err)
		}
		debtTable.Debt = debt
		debtTable.Date = date
		debtsInfo = append(debtsInfo, debtTable)
	}
	defer db.Close()
	return debtsInfo
}

func curdate() interface{} {
	dt := time.Now()
	//Format MM-DD-YYYY
	dtFMT := dt.Format("01-02-2006")
	return dtFMT

}
