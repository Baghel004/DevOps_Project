package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp" // NEW
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", homePage)
	mux.HandleFunc("/courses", coursePage)
	mux.HandleFunc("/about", aboutPage)
	mux.HandleFunc("/contact", contactPage)


	http.Handle("/metrics", promhttp.Handler()) // NEW
	mux.Handle("/metrics", ExposePrometheusHandler())
	instrumented := Instrument(mux)

	err := http.ListenAndServe("0.0.0.0:9090", instrumented)
	if err != nil {
		log.Fatal(err)
	}
}
