package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// FileData represents the data passed to the template
type FileData struct {
	Title   string
	Files   []string
	Message string
}

// HomeHandler serves the home page by rendering the template with dynamic data
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template from the web/templates directory
	tmpl, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		log.Printf("template parsing error: %v", err)
		return
	}

	// Example data passed to the template
	data := FileData{
		Title:   "My Dynamic Go Webpage",
		Files:   []string{"file1.txt", "file2.jpg", "file3.pdf"}, // Dynamic list of files
		Message: "Welcome to the dynamically rendered page!",
	}

	// Set the content type as HTML
	w.Header().Set("Content-Type", "text/html")

	// Execute the template and pass the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("template execution error: %v", err)
	}
}
