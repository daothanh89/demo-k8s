package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/quotes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			return
		}

		resp, err := http.Get(os.Getenv("QUOTE_SERVICE_ENDPOINT"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			print(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	http.Handle("/", http.FileServer(http.Dir("static/build")))

	log.Println("[INFO] Server running on port: 8000")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), logger(http.DefaultServeMux)))
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Host, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
