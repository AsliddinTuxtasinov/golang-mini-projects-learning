package handler

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title    string `json:"title"`
	FullName string `json:"full_name"`
	Header   string `json:"header"`
	Age      int    `json:"age"`
	Content  string `json:"content"`
	URL      string `json:"url"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// body, _ := utils.LoadFile("index.html")
	// fmt.Fprintf(w, body)

	page := Page{
		Title:    "TITLE",
		FullName: "ASLIDDIN TUKHTASINOV",
		Header:   "HEADER",
		Age:      23,
		Content:  "CONTENT",
		URL:      "https://www.practical-go-lessons.com/chap-3-the-terminal",
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, page)
}
