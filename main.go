package main

import "net/http"

func main() {
	http.HandleFunc("/calc", getWasteCalc)
	http.ListenAndServe(":8080", nil)
}
