package service

import (
	"HouseCalculator/repo"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"math"
	"net/http"
	"strconv"
)

type UserLogin struct {
	UserEmail string `json:"email"`
	UserID    string `json:"id"`
}

type Debt struct {
	Debt1        string `json:"Debt1"`
	Debt2        string `json:"Debt2"`
	Debt3        string `json:"Debt3"`
	Debt4        string `json:"Debt4"`
	Debt5        string `json:"Debt5"`
	Debt6        string `json:"Debt6"`
	Debt7        string `json:"Debt7"`
	Debt8        string `json:"Debt8"`
	Debt9        string `json:"Debt9"`
	Debt10       string `json:"Debt10"`
	TotalDebt    string `json:"total_debt"`
	YearlyIncome string `json:"yearly_income"`
	DTI          string `json:"dti"`
}

var UserList []UserLogin

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "528673173486-np42b4f1opj6u0s2akgiskd2iuj5fqh4.apps.googleusercontent.com",
		ClientSecret: "fv4OfNOfMC7uJXHuxnQxcgD6",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	//TODO: randomize it
	randomState  = "random"
	loggedInUser UserLogin
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><h1 style="text-align:center">Please Sign in With</h1><body style="text-align:center"><a href="/login">Google Log In</a></body></html>`
	fmt.Fprintf(w, html)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallBack(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	println(r.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create request: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	//User := new(UserLogin)
	println("######")
	err = json.NewDecoder(resp.Body).Decode(&loggedInUser)
	if err != nil {
		fmt.Printf("could not decode response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	println(loggedInUser.UserEmail)
	println("######")

	defer resp.Body.Close()
	//repo.DBUserInsert(User.UserID, User.UserEmail)

	http.Redirect(w, r, "/assets/form.html", http.StatusFound)

	return
}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/assets/login.html", http.StatusFound)
}

/*func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
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

}*/

var debtList []Debt

func GetDebtMarshalTotal(w http.ResponseWriter, r *http.Request) {
	println("getDEBT")

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
	CheckErr(err)
}

func CreateDebtCalculationHandler(w http.ResponseWriter, r *http.Request) {
	println("postDEBT")
	// Create a new instance of Debt

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
	//GetDebts(r)
	debt := GetDebts(r)

	// Append our existing list of Debts with a new entry

	intDebt1, intDebt2, intDebt3, intDebt4, intDebt5, intDebt6, intDebt7, intDebt8, intDebt9, intDebt10 := DebtsToFloat(debt)

	CombinedTotalDebt := TotalDebt(intDebt1, intDebt2, intDebt3, intDebt4, intDebt5, intDebt6, intDebt7, intDebt8, intDebt9, intDebt10)
	debt.TotalDebt = strconv.Itoa(int(CombinedTotalDebt))
	calcDTI := CalculatesDTI(debt.TotalDebt, debt.YearlyIncome)
	debt.DTI = calcDTI
	repo.DBDebtInsert(debt.TotalDebt, debt.DTI, loggedInUser.UserID, loggedInUser.UserEmail)
	//debtList is what's being called by the front end (Get /debt handler)
	debtList = append(debtList, debt)
	http.Redirect(w, r, "/assets/form2.html", http.StatusFound)

}

func DebtsToFloat(debt Debt) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	intDebt1, _ := strconv.ParseFloat(debt.Debt1, 64)
	intDebt2, _ := strconv.ParseFloat(debt.Debt2, 64)
	intDebt3, _ := strconv.ParseFloat(debt.Debt3, 64)
	intDebt4, _ := strconv.ParseFloat(debt.Debt4, 64)
	intDebt5, _ := strconv.ParseFloat(debt.Debt5, 64)
	intDebt6, _ := strconv.ParseFloat(debt.Debt6, 64)
	intDebt7, _ := strconv.ParseFloat(debt.Debt7, 64)
	intDebt8, _ := strconv.ParseFloat(debt.Debt8, 64)
	intDebt9, _ := strconv.ParseFloat(debt.Debt9, 64)
	intDebt10, _ := strconv.ParseFloat(debt.Debt10, 64)
	return intDebt1, intDebt2, intDebt3, intDebt4, intDebt5, intDebt6, intDebt7, intDebt8, intDebt9, intDebt10
}

func GetDebts(r *http.Request) Debt {
	debt := Debt{}
	debt.YearlyIncome = r.Form.Get("yearly_income")
	debt.Debt1 = r.Form.Get("Debt1")
	debt.Debt2 = r.Form.Get("Debt2")
	debt.Debt3 = r.Form.Get("Debt3")
	debt.Debt4 = r.Form.Get("Debt4")
	debt.Debt5 = r.Form.Get("Debt5")
	debt.Debt6 = r.Form.Get("Debt6")
	debt.Debt7 = r.Form.Get("Debt7")
	debt.Debt8 = r.Form.Get("Debt8")
	debt.Debt9 = r.Form.Get("Debt9")
	debt.Debt10 = r.Form.Get("Debt10")
	return debt
}

func TotalDebt(intDebt1 float64, intDebt2 float64, intDebt3 float64, intDebt4 float64, intDebt5 float64, intDebt6 float64, intDebt7 float64, intDebt8 float64, intDebt9 float64, intDebt10 float64) float64 {
	TotalDebt := intDebt1 + intDebt2 + intDebt3 + intDebt4 + intDebt5 + intDebt6 + intDebt7 + intDebt8 + intDebt9 + intDebt10
	return TotalDebt
}

/*func GetHistoryHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("debtHistory.html"))
	table := repo.DbSelectQuery()
	err := tmpl.ExecuteTemplate(w, "debtHistory", table)
	CheckErr(err)
}*/

func UserHistoryHandler(w http.ResponseWriter, r *http.Request) {
	table := repo.DbSelectQuery(loggedInUser.UserEmail)
	for i := range table {
		emp := table[i]
		fmt.Fprintf(w, "%v|%v|%v|%v|\n", emp.Debt, emp.Date, emp.DTI, emp.UserEmail)
	}
}

func CalculatesDTI(td string, yi string) string {
	convertedTD, _ := strconv.ParseFloat(td, 64)
	println(convertedTD)
	convertedYI, _ := strconv.ParseFloat(yi, 64)
	monthlyYI := convertedYI / 12
	dti := convertedTD / monthlyYI
	dtiPercent := dti * 100
	wholeValueDTI := math.Round(dtiPercent)
	strConvertedValue := strconv.FormatFloat(wholeValueDTI, 'f', 0, 64)
	return strConvertedValue
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
