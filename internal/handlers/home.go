package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// HomeHandler serves the home page by rendering the template
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template from the web/templates directory
	tmpl, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		log.Printf("template parsing error: %v", err)
		return
	}

	// Set the content type as HTML
	w.Header().Set("Content-Type", "text/html")

	// Execute the template and render it
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Printf("template execution error: %v", err)
	}
}
