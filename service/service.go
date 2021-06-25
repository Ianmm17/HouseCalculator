package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserLogin struct {
	UserEmail    string `json:"username"`
	UserPassword string `json:"password"`
}

type Debt struct {
	Debt1     string `json:"Debt1"`
	Debt2     string `json:"Debt2"`
	Debt3     string `json:"Debt3"`
	Debt4     string `json:"Debt4"`
	Debt5     string `json:"Debt5"`
	TotalDebt string `json:"total_debt"`
}

var UserList []UserLogin

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/assets/login.html", http.StatusFound)
	/*UserListBytes, err := json.Marshal(UserList)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(UserListBytes)
	if err != nil {
		return
	}*/
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	user := UserLogin{}
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.UserEmail = r.Form.Get("username")
	user.UserPassword = r.Form.Get("password")

	UserList = append(UserList, user)
	fmt.Println(UserList)
	if user.UserEmail == "milleri" && user.UserPassword == "password123" {
		http.SetCookie(w, &http.Cookie{
			Name:  "isLoggedIn",
			Value: "success",
		})

		http.Redirect(w, r, "/assets/form.html", http.StatusFound)
	} else {
		http.Redirect(w, r, "/assets/login.html", http.StatusFound)
	}

}

var debtList []Debt

func GetDebtHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("isLoggedIn")
	if err != nil {
		println("made it here")
		http.Redirect(w, r, "/assets/login.html", http.StatusFound)
	} else {
		c = &http.Cookie{
			Name: "isLoggedIn"}
		http.SetCookie(w, c)

		http.Redirect(w, r, "/assets/form.html", http.StatusFound)

	}
	//Convert the "debList" variable to json
	DebtListBytes, err := json.Marshal(debtList)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of Debts to the response
	_, err = w.Write(DebtListBytes)
	if err != nil {
		return
	}
}

func CreateDebtHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Debt
	debt := Debt{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the Debts from the form info
	debt.Debt1 = r.Form.Get("Debt1")
	debt.Debt2 = r.Form.Get("Debt2")
	debt.Debt3 = r.Form.Get("Debt3")
	debt.Debt4 = r.Form.Get("Debt4")
	debt.Debt5 = r.Form.Get("Debt5")

	// Append our existing list of Debts with a new entry

	intDebt1, _ := strconv.ParseFloat(debt.Debt1, 64)
	intDebt2, _ := strconv.ParseFloat(debt.Debt2, 64)
	intDebt3, _ := strconv.ParseFloat(debt.Debt3, 64)
	intDebt4, _ := strconv.ParseFloat(debt.Debt4, 64)
	intDebt5, _ := strconv.ParseFloat(debt.Debt5, 64)
	totalDebt := TotalDebt(intDebt1, intDebt2, intDebt3, intDebt4, intDebt5)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	debt.TotalDebt = strconv.Itoa(int(totalDebt))
	debtList = append(debtList, debt)
	//http.Redirect(w, r, "/assets/form.html", http.StatusFound)
	fmt.Fprintf(w, "%.2f\n", totalDebt)

}

func TotalDebt(intDebt1 float64, intDebt2 float64, intDebt3 float64, intDebt4 float64, intDebt5 float64) float64 {
	TotalDebt := intDebt1 + intDebt2 + intDebt3 + intDebt4 + intDebt5
	return TotalDebt
}
