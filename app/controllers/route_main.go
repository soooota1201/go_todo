package controllers

import (
	// "log"
	"net/http"
	// "text/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
}