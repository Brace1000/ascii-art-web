package handler

import (

	"html/template"
	"net/http"
	"strings"

	"piscine/ascii"
)

type Data struct {
	Text   string
	Result string
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

// function to get the home page for the users
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "HTTP status 404 - page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		// Parse the template file
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Execute the template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// function for reading the text from the user and converting it to ascii
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		font := r.FormValue("file")

		for _, char := range text {
			if char > '~' {
				http.Error(w, "400 Bad Request: Unrecognized ascii chatacters", http.StatusBadRequest)
				return
			}
		}

		if text == "" || font == "" {
			http.Error(w, "400 Bad Request: Missing required fields", http.StatusBadRequest)
			return
		}

		var print string

		if strings.Contains(text, "\n") {
			myslice := strings.Split(text, "\n")
			for _, line := range myslice {
				print += ascii.PrintAsci(font, line)
			}
		} else {
			print = ascii.PrintAsci(font, text)
			if print == "nil" {
				w.WriteHeader(http.StatusNotFound)
				tpl, _ := template.ParseFiles("templates/404.html")
				tpl.Execute(w, nil)
				return
			}
		}
		data := Data{
			Text:   text,
			Result: print,
		}
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		tmpl.Execute(w, data)
	} else if r.Method != http.MethodPost {
		http.Error(w, "HTTP status 405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
