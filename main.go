// Playing with Go... January 2022
// Tutorial on Youtube www.youtube.com/watch?v=0sRjYzL_oYs
// To use:
//   1. start the webserver in the terminal go run main.go
//   2. from Chrome, go to the server localhost:8070/aboutme

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Handler for "/"
func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP handler for / called...")
	fmt.Fprintf(w, "This is the root dir, hello everyone!")
}

// Handler for "/aboutme"
func handleAboutMe(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP handler for /aboutme...")
	// fmt.Printf("%v\n\n", r)

	fmt.Fprintf(w, "You have called /aboutme, and I am JP Chauvet")
}

// Start app webserver
func main() {
	log.Println("Starting webserver...")

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/aboutme", handleAboutMe)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the webserver
	log.Printf("Webserver listening on port %s", port)
	http.ListenAndServe(":"+port, nil)

}
