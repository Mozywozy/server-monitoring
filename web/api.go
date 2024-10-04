package web

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"server-monitor/internal/storage"
	"path/filepath"
)

// PageData menyimpan data yang akan dikirim ke template
type PageData struct {
	Servers []ServerStatus
}

// ServerStatus menyimpan status server
type ServerStatus struct {
	Name   string
	Status string
}

func StartAPIServer() {
	http.HandleFunc("/", RedirectToStatus)
	http.HandleFunc("/status", StatusHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// RedirectToStatus mengarahkan ke halaman status
func RedirectToStatus(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/status", http.StatusSeeOther)
}

// StatusHandler menampilkan status server
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var serverStatuses []ServerStatus

	// Mengambil data status server dari database
	rows, err := storage.GetAllServerStatus()
	if err != nil {
		http.Error(w, "Unable to fetch server status", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var status ServerStatus
		if err := rows.Scan(&status.Name, &status.Status); err != nil {
			http.Error(w, "Error reading row", http.StatusInternalServerError)
			return
		}
		serverStatuses = append(serverStatuses, status)
	}

	// Membangun path untuk template secara dinamis
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	templatePath := filepath.Join(currentDir, "web", "static", "status.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Error loading template:", err) // Tambahkan log untuk error
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	pageData := PageData{Servers: serverStatuses}
	tmpl.Execute(w, pageData)
}
