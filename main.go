package main

import (
	"fmt"
	"time"
	"net/http"
)

func main() {
	fmt.Println("Starting web server...")
	http.HandleFunc("/", handler)
    http.ListenAndServe(":1080", nil)
}

// Show calendar on GET and 501 error on any other requests
func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
    case "GET":     
        calendar(w)
    default:
    	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
    }
}

// Show the current month
func calendar(w http.ResponseWriter) {
	now := time.Now()
	today := now.Day()
	month := now.Month()
	year := now.Year()
	location := now.Location()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1).Day()

	fmt.Fprintf(w, "%v %v\n", month, year)
	fmt.Fprintf(w, "----------\n")
	for i := 1; i <= lastOfMonth; i++ {
		day := time.Date(year, month, i, 0, 0, 0, 0, location)
		if i == today {
			fmt.Fprintf(w, "%v %v (today)\n", i, day.Weekday())
		} else {
			fmt.Fprintf(w, "%v %v\n", i, day.Weekday())
		}
	}
}
