package main

import (
	"HouseCalculator/service"
	"github.com/gorilla/mux"
	"net/http"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/login.html" in our browser, the file server
	// will look for only "login.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/login.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	//
	r.HandleFunc("/", service.HandleHome)
	r.HandleFunc("/login", service.HandleLogin)
	r.HandleFunc("/callback", service.HandleCallBack)
	r.HandleFunc("/debt", service.GetDebtMarshalTotal).Methods("GET")
	r.HandleFunc("/debt", service.CreateDebtCalculationHandler).Methods("POST")
	r.HandleFunc("/history", service.UserHistoryHandler).Methods("POST")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}
