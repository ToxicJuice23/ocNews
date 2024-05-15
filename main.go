package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type alertInfo struct {
	Title         string `json:"Title"`
	RouteNumber   string `json:"RouteNumber"`
	DateEffective string `json:"DateEffective"`
	Desc          string `json:"Description"`
}

func getAlerts() []alertInfo {
	r := getFromUrl("https://www.octranspo.com/en/alerts", time.Second*20)
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		_ = fmt.Errorf("%v", err.Error())
		os.Exit(1)
	}

	var alerts []alertInfo
	doc.Find(".alert").Each(func(i int, s *goquery.Selection) {
		route, exists := s.Attr("data-routes")
		if exists {
			alert := alertInfo{}
			alert.Title = s.Find(".accordion").First().Text()
			alert.RouteNumber = route
			fmt.Printf("rn: %s\n", route)
			infos := []string{}
			content := s.Find(".accordion-content")
			content = content.First().Children()

			content.Each(func(i int, se *goquery.Selection) {
				infos = append(infos, se.Text())
			})
			alert.DateEffective = infos[1]
			alert.Desc = infos[2]
			alerts = append(alerts, alert)
		}
	})
	return alerts
}

func return500(w http.ResponseWriter) {
	w.WriteHeader(500)
	fmt.Fprintf(w, "Internal Server Error")
}

func return404(w http.ResponseWriter) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "File not found")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	alerts := getAlerts()
	encoder := json.NewEncoder(w)
	err := encoder.Encode(alerts)
	if err != nil {
		return500(w)
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	res := r.URL.Path
	fmt.Fprintf(os.Stdout, "Recieved request: %s\n", time.Now())
	fmt.Fprintf(os.Stdout, "Path: %v\n", res)
	var err error
	var buf []byte
	if res == "/" {
		buf, err = os.ReadFile("./public/index.html")
		fmt.Printf("sent home\n")
	} else {
		buf, err = os.ReadFile("./public/" + res)
	}
	if err != nil {
		_ = fmt.Errorf("%v", err.Error())
		return404(w)
	}
	w.Write(buf)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodHead {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Illegal Method, must be HEAD.\nRecieved method: %s\n", r.Method)
	} else {
		w.WriteHeader(200)
	}
}

func main() {
	fmt.Fprintf(os.Stdout, "Initializing server.\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api", apiHandler)
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
