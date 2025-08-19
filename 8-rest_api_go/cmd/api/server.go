// SPDX-License-Identifier: MIT
package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

"restapi/internal/repository/sqlconnect"
)

type Teacher struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
}

var dbConn = initDB()

func initDB() *sql.DB {
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Fatal("❌ Error connecting to DB: ", err)
	}
	fmt.Println("✅ Connected to MariaDB/MySQL")
	return db
}

// ------------------ Handlers ------------------

// Create Teacher
func createTeacherHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var t Teacher
	if err := json.Unmarshal(body, &t); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := dbConn.Exec(
		"INSERT INTO teachers (first_name, last_name, class, subject) VALUES (?, ?, ?, ?)",
		t.FirstName, t.LastName, t.Class, t.Subject,
	)
	if err != nil {
		http.Error(w, "Error inserting teacher", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	t.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Teacher created successfully",
		"teacher": t,
	})
}

// Get All Teachers
func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := dbConn.Query("SELECT id, first_name, last_name, class, subject FROM teachers")
	if err != nil {
		http.Error(w, "Error fetching teachers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	teachers := []Teacher{}
	for rows.Next() {
		var t Teacher
		if err := rows.Scan(&t.ID, &t.FirstName, &t.LastName, &t.Class, &t.Subject); err != nil {
			http.Error(w, "Error scanning teacher", http.StatusInternalServerError)
			return
		}
		teachers = append(teachers, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

// Update Teacher by ID
func updateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var t Teacher
	if err := json.Unmarshal(body, &t); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err = dbConn.Exec(
		"UPDATE teachers SET first_name=?, last_name=?, class=?, subject=? WHERE id=?",
		t.FirstName, t.LastName, t.Class, t.Subject, id,
	)
	if err != nil {
		http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return
	}

	t.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Teacher updated successfully",
		"teacher": t,
	})
}

// Delete Teacher by ID
func deleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = dbConn.Exec("DELETE FROM teachers WHERE id=?", id)
	if err != nil {
		http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": fmt.Sprintf("Teacher with ID %d deleted successfully", id),
	})
}

// ------------------ Server ------------------

func main() {
	defer dbConn.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTeachersHandler(w, r)
		case http.MethodPost:
			createTeacherHandler(w, r)
		case http.MethodPut:
			updateTeacherHandler(w, r)
		case http.MethodDelete:
			deleteTeacherHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	fmt.Println("🚀 HTTPS server started on https://localhost:3000")
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
