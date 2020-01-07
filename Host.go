package main

import (
	"fmt"
	"net/http"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fucking AungThu! Showing My Penic.Can See?")
}
func indesName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hus Lae Yi Hnin")
}
func mamaHnin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hnin Thu Aung")
}
func main() {
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/name/", indesName)
	http.HandleFunc("/hnin", mamaHnin)
	http.ListenAndServe(":8002", nil)
}
