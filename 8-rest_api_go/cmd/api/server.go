// SPDX-License-Identifier: MIT
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// -------------------- Struct --------------------
type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

// -------------------- Handlers --------------------
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
}

// Teachers Handler
func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Teachers Route"))
		fmt.Println("Hello GET Method on Teachers Route")

	case http.MethodPost:
		// Read RAW Body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		fmt.Println("RAW Body:", string(body))

		// Parse JSON
		var userInstance user
		err = json.Unmarshal(body, &userInstance)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Println("Parsed JSON:", userInstance)

		// Respond back in JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Teacher created successfully",
			"user":    userInstance,
		})

	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
		fmt.Println("Hello PUT Method on Teachers Route")

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
		fmt.Println("Hello PATCH Method on Teachers Route")

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
		fmt.Println("Hello DELETE Method on Teachers Route")
	}
}

// Students Handler
func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Students Route"))
		fmt.Println("Hello GET Method on Students Route")

	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Students Route"))
		fmt.Println("Hello POST Method on Students Route")

	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Students Route"))
		fmt.Println("Hello PUT Method on Students Route")

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Students Route"))
		fmt.Println("Hello PATCH Method on Students Route")

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Students Route"))
		fmt.Println("Hello DELETE Method on Students Route")
	}
}

// Execs Handler (dummy example)
func execsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Execs Route"))
	fmt.Println("Hello Execs Route")
}

// -------------------- Main --------------------
func main() {
	port := ":3000"
	cert := "cert.pem"
	key := "key.pem"

	// Create a new ServeMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	// TLS configuration
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create custom HTTPS server
	server := &http.Server{
		Addr:      port,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("✅ Server is running on https://localhost" + port)

	// Start HTTPS server with TLS cert & key
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("❌ Error starting the TLS server:", err)
	}
}
