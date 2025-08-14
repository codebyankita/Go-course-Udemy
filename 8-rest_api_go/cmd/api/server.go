package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Struct for JSON parsing
type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("----- Incoming Request -----")
		fmt.Println("Method:", r.Method)

		// Parse form data (necessary for x-www-form-urlencoded)
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		fmt.Println("Form:", r.Form)

		// Prepare response data from form
		response := make(map[string]interface{})
		for key, value := range r.Form {
			response[key] = value[0]
		}
		if len(response) > 0 {
			fmt.Println("Processed Form Map:", response)
		}

		// --- RAW Body parsing ---
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		fmt.Println("RAW Body (string):", string(body))

		// If expecting JSON -> unmarshal into struct
		var userInstance User
		if len(body) > 0 {
			err = json.Unmarshal(body, &userInstance)
			if err == nil {
				fmt.Println("Unmarshaled JSON into User struct:", userInstance)
				fmt.Println("Received user name as:", userInstance.Name)
			} else {
				fmt.Println("Error unmarshaling into struct:", err)
			}
		}

		// Unmarshal JSON into map as well
		response1 := make(map[string]interface{})
		if len(body) > 0 {
			err = json.Unmarshal(body, &response1)
			if err == nil {
				fmt.Println("Unmarshaled JSON into a map:", response1)
			} else {
				fmt.Println("Error unmarshaling into map:", err)
			}
		}

		// Log request details
		fmt.Println("Request Details:")
		fmt.Println("Form:", r.Form)
		fmt.Println("Header:", r.Header)
		fmt.Println("ContentLength:", r.ContentLength)
		fmt.Println("Host:", r.Host)
		fmt.Println("Protocol:", r.Proto)
		fmt.Println("RemoteAddr:", r.RemoteAddr)
		fmt.Println("RequestURI:", r.RequestURI)
		fmt.Println("TLS:", r.TLS)
		fmt.Println("Trailer:", r.Trailer)
		fmt.Println("TransferEncoding:", r.TransferEncoding)
		fmt.Println("URL:", r.URL)
		fmt.Println("UserAgent:", r.UserAgent())
		fmt.Println("Port:", r.URL.Port())
		fmt.Println("URL Scheme:", r.URL.Scheme)

		fmt.Println("-----------------------------")

		// Respond with JSON
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(map[string]string{"message": "Hello GET Method on Teachers Route"})

		case http.MethodPost:
			if len(response1) > 0 {
				json.NewEncoder(w).Encode(response1)
			} else if (userInstance != User{}) {
				json.NewEncoder(w).Encode(userInstance)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"message": "Hello POST Method on Teachers Route"})
			}

		case http.MethodPut:
			json.NewEncoder(w).Encode(map[string]string{"message": "Hello PUT Method on Teachers Route"})

		case http.MethodPatch:
			json.NewEncoder(w).Encode(map[string]string{"message": "Hello PATCH Method on Teachers Route"})

		case http.MethodDelete:
			json.NewEncoder(w).Encode(map[string]string{"message": "Hello DELETE Method on Teachers Route"})

		default:
			http.Error(w, `{"error":"Method Not Allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello execs Route"))
		fmt.Println("Hello execs Route")
	})

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
