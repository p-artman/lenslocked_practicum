package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
	// tplPath := filepath.Join("templates", "faq.gohtml")
	// executeTemplate(w, tplPath)
}

func galleriesHandler(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "id") != "" {
		galleryID := chi.URLParam(r, "id")
		w.Write([]byte(fmt.Sprintf("<p>gallery id is %v</p>", galleryID)))
	} else {
		fmt.Fprint(w, "<p>the list of all galleries</p>")
	}
}

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("processing template: %v", err)
		http.Error(w, "There was an error processing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.RequestID)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.Route("/galleries", func(r chi.Router) {
		r.Get("/", galleriesHandler)
		r.Get("/{id}", galleriesHandler)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
	log.Fatal(http.ListenAndServe(":3000", r))
	// Stopping the Host Network Service on Windows (net stop hns) solved the problem.
}
