package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init(){
	tmpl=template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "template1.html",nil)
}

func selectionHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "template2.html",nil)
}

func rulesHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "template3.html",nil)
}

func main() {
	fs:=http.FileServer(http.Dir("assets"))
	http.Handle("/assets/",http.StripPrefix("/assets", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/template2.html", selectionHandler)
	http.HandleFunc("/template3.html", rulesHandler)
	http.ListenAndServe(":9999", nil)
}