package pkg

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartServer() {
	h := handlers{users: map[string]string{}}
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/api/login", h.login)
	r.HandleFunc("/api/register", h.register)
	r.HandleFunc("/api/quotes", authMiddleware(h.quoteGet))

	r.PathPrefix("/").Handler(spaHandler{
		staticPath: "static/build",
		indexPath:  "index.html",
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("[INFO] Server running on port: 8000")
	log.Fatal(srv.ListenAndServe())
}
