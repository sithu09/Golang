package main

import (
	"fmt"
	"net/http"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fucking CDE! I Showing My Penic.Can See?")
}
func indesName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hus Lae Yi Hnin")
}
func main() {
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/name/", indesName)
	http.ListenAndServe(":8001", nil)
}
