package main // Entry point package. Every executable Go app must use package main.//Folder name == Package name

import (
	"log"      // Used for logging messages and errors
	"net/http" // Provides HTTP server and routing functionality
	"crud-go/handlers"
	"strings"
)

func main() {

	// Register a route "/hello"
	// When a request hits this path, the handler function runs.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		// Send response back to client.
		// HTTP responses are sent as bytes, so convert string → []byte.
		w.Write([]byte("Server running..."))
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request){
		// EXACT match → /users
		if r.URL.Path == "/users/" {

			if r.Method == http.MethodGet {
				handlers.GetUsers(w, r)
				return
			}

			if r.Method == http.MethodPost {
				handlers.CreateUser(w, r)
				return
			}
		}
		// PREFIX match → /users/{id}
		if strings.HasPrefix(r.URL.Path, "/users/") {

			if r.Method == http.MethodGet {
				handlers.GetUsersById(w, r)
				return
			}

			if r.Method == http.MethodPut {
				handlers.UpdateUser(w, r)
				return
			}

			if r.Method == http.MethodDelete {
				handlers.DeleteUser(w, r)
				return
			}
		}
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
	})

	// Log message to indicate server startup
	log.Println("Server started at :8080")

	// Start HTTP server on port 8080
	// nil → use default router (ServeMux)
	// log.Fatal logs error if server fails and stops the program
	log.Fatal(http.ListenAndServe(":8080", nil))
}


