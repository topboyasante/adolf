package templates

func GenerateModelTemplate() string {
	modelTemplate := `
	package models
		
	import (
		"time"
	)
		
	// Define a Model type to hold the data for an individual Model. 
	type Model struct {
		ID      int
		Title   string
		Content string
		Created time.Time
		Expires time.Time
	}
	`
	return modelTemplate
}

func GenerateRouteTemplate() string {
	modelTemplate := `
	package main

	import "net/http"
	
	// The routes() method returns a servemux containing our application routes.
	func routes() *http.ServeMux {
		mux := http.NewServeMux()
	
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
		
		return mux
	}
	`
	return modelTemplate
}
