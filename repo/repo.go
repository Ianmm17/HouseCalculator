package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type UserDebtInfo struct {
	Debt      string
	Date      string
	UserEmail string
	UserID    string
	DTI       string
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

func DBUserInsert(userEmail string, userID string) {
	//DataBase Inserting debt and date
	db := DBSetup()
	stmt, err := db.Prepare("INSERT INTO DebtStorage.DebtTable (UserEmail, UserID) VALUES (?, ?)")
	if err != nil {
		log.Printf("Error %s when it is preparing", err)
	}
	_, err = stmt.Exec(userEmail, userID)
	if err != nil {
		log.Printf("Error %s when inserting DB\n", err)
	}
}

func DBDebtInsert(debt string, dti string, userID string, userEmail string) {
	//DataBase Inserting debt and date
	db := DBSetup()
	dti = dti + "%"
	stmt, err := db.Prepare("INSERT INTO DebtStorage.DebtTable (Debt, Date, DTI, UserID, UserEmail) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("Error %s when it is preparing", err)
	}
	_, err = stmt.Exec(debt, curdate(), dti, userID, userEmail)
	if err != nil {
		log.Printf("Error %s when inserting DB\n", err)
	}
}

func DbSelectQuery(userEmail string) []UserDebtInfo {
	db := DBSetup()
	stmt, err := db.Prepare(`SELECT UserEmail, Debt, DTI, Date FROM DebtStorage.DebtTable WHERE UserEmail = ?`)
	if err != nil {
		log.Println(err)
	}
	rows, err := stmt.Query(userEmail)
	if err != nil {
		log.Printf("Error %s when inserting DB\n", err)
	}
	debtTable := UserDebtInfo{}
	debtsInfo := []UserDebtInfo{}

	for rows.Next() {
		var debt, date, dti, userEmail string
		err = rows.Scan(&debt, &date, &dti, &userEmail)
		if err != nil {
			log.Println(err)
		}
		debtTable.Debt = debt
		debtTable.Date = date
		debtTable.UserEmail = userEmail
		debtTable.DTI = dti
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
