package handler

import (
	"html/template"
	"net/http"
)

type User struct {
	Name              string
	PrimaryProgram    string
	AmateurishProgram string
	Highlight         bool
}

func Index(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{"张三", "PHP", "Go", true},
		{"李四", "Java", "Go", false},
		{"王五", "Go", "Rust", true},
	}
	t, _ := template.ParseFiles("templates/index/index.html")
	t.ExecuteTemplate(w, "index/index.html", users)
}
