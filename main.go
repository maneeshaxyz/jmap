package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/get/resource", getHandler)
	mux.HandleFunc("/post/resource", postHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from your JMAP server"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my GET handler"))
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Add("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This is my POST handler"))
}

// func main() {
// 	server := &JMAPServer{}

// 	srv := &http.Server{
// 		Addr:         ":8080",
// 		Handler:      server,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  120 * time.Second,
// 	}

// 	log.Println("JMAP server starting on port 8080")
// 	log.Fatal(srv.ListenAndServe())
// }
