package xss

import (
	htmlTemplate "html/template"
	"net/http"
	"text/template"
)

func XSSHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/xss/xss.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "xss.html", struct {
		Path string
	}{Path: "» xss"})
}
func Easy(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/xss/xss_easy.html", "templates/base.html"))
	payload := r.FormValue("payload")

	tmpl.ExecuteTemplate(w, "xss_easy.html", struct {
		Payload string
	}{Payload: payload})

}

func Hard(w http.ResponseWriter, r *http.Request) {
	tmpl := htmlTemplate.Must(htmlTemplate.ParseFiles("templates/xss/xss_hard.html", "templates/base.html"))
	payload := r.FormValue("payload")

	tmpl.ExecuteTemplate(w, "xss_hard.html", struct {
		Payload string
	}{Payload: payload})

}
