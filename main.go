package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"html/template"
	"net/http"
	"os"
	"time"
)

func getAlerts() []string {
	r := curl("https://www.octranspo.com/en/alerts", time.Second*20)
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		fmt.Errorf("Error: %v\n", err)
		os.Exit(1)
	}

	var alerts []string
	doc.Find(".alert").Find(".accordion").Each(func(i int, s *goquery.Selection) {
		alerts = append(alerts, s.Text())
	})
	return alerts
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "Recieved request: %s\n", time.Now())
	fd, err := os.ReadFile("template.html")
	t, err := template.New("template").Parse(string(fd))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal Server Error")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	t.Execute(w, getAlerts())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" || r.Method != "head" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Illegal Method, must be HEAD.\n")
	} else {
		w.WriteHeader(200)
	}
}

func main() {
	fmt.Fprintf(os.Stdout, "Initializing server.\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":80"
	} else {
		port = ":" + port
	}
	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(port, mux)
}
