package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wisnu-bdd/be-go-with-auth-template/config"
	"github.com/wisnu-bdd/be-go-with-auth-template/db"
	"github.com/wisnu-bdd/be-go-with-auth-template/handlers"
	"github.com/wisnu-bdd/be-go-with-auth-template/middleware"
)

func main() {
	// Load environment configuration
	config.Load()

	if err := db.ConnectToMongo(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Define routes with middleware
	http.HandleFunc("/", middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend Go with auth template!")
	}))

	http.HandleFunc("/user/register", middleware.CORS(handlers.Register))
	http.HandleFunc("/user/login", middleware.CORS(handlers.Login))
	http.HandleFunc("/user/get-me", middleware.CORS(middleware.JWT(handlers.GetMe)))
	http.HandleFunc("/user/get-all", middleware.CORS(middleware.JWT(handlers.GetUsers)))
	http.HandleFunc("/user/update-details", middleware.CORS(middleware.JWT(handlers.UpdateUserDetailsByEmail)))
	http.HandleFunc("/user/update-password", middleware.CORS(middleware.JWT(handlers.UpdateUserPasswordByEmail)))

	http.HandleFunc("/protected", middleware.CORS(middleware.JWT(handlers.Protected)))

	// Start server
	port := config.Port
	fmt.Printf("Server running on http://localhost:%s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  POST /login     - Login with email/password")
	fmt.Println("  GET  /protected - Protected endpoint (requires JWT)")

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
