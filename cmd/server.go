package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	mw "gocourse/cmd/internal/middlewares"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		fmt.Println(r.URL.Path)

		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")
		fmt.Println("The ID is :", userID)

		fmt.Println("Query params", r.URL.Query())
		queryParams := r.URL.Query()

		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")

		if sortorder == "" {
			sortorder = "DESC"
		}

		fmt.Printf("SortBy : %v, SortOrder: %v, key: %v", sortby, sortorder, key)

		w.Write([]byte("Hello GET METHOD on Teachers Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST METHOD on Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH METHOD on Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE METHOD on Teachers Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT METHOD on Teachers Route"))
	}

}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET METHOD on Students Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST METHOD on Students Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH METHOD on Students Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE METHOD on Students Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT METHOD on Students Route"))
	}

}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET METHOD on Execs Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST METHOD on Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH METHOD on Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE METHOD on Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT METHOD on Execs Route"))
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	// ! 1. SCALLABLE API
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	server := &http.Server{
		Addr: port,
		// Handler:   mux,
		Handler:   rl.Middleware((mw.Compression(mw.ResponseTimeMiddlewares(mw.SecurityHeaders(mw.Cors(mux)))))),
		TLSConfig: tlsConfig,
	}

	// ! 2. SMALL API
	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/teachers/", teachersHandler)
	// http.HandleFunc("/students/", studentsHandler)
	// http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running in port", port)
	// err := http.ListenAndServe(port, nil)
	err := server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
