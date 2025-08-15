// SPDX-License-Identifier: MIT
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	mw "restapi/internal/api/middlewares"
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

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Teachers Route"))

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var userInstance user
		err = json.Unmarshal(body, &userInstance)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Teacher created successfully",
			"user":    userInstance,
		})

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Students Route"))
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Execs Route"))
}

// -------------------- Main --------------------
func main() {
	port := ":3000"
	cert := "cert.pem"
	key := "key.pem"

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	// Wrap middlewares in correct order
	handler := mw.ResponseTimeMiddleware(
		mw.SecurityHeaders(
			mw.Cors(mux),
		),
	)

	// TLS Config
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	server := &http.Server{
		Addr:      port,
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	fmt.Println("✅ Server is running on https://localhost" + port)

	if err := server.ListenAndServeTLS(cert, key); err != nil {
		log.Fatalln("❌ Error starting the TLS server:", err)
	}
}
