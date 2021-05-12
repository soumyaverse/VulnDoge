package xss

import (
	htmlTemplate "html/template"
	"net/http"
	"text/template"
)

func XSSHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/xss/xss.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "xss.html", nil)
}
func Easy(w http.ResponseWriter, r *http.Request) {
	payload := r.FormValue("payload")

	tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "easy1.html", struct {
		Title   string
		Payload string
		Desc    string
		Login   bool
		User    string
	}{Title: "xss easy", Desc: `<h2>XSS</h2>
	` + payload + `
	<br>
	<form action="/xss/easy/" method="get">
		<input type="text" name="payload"><br><br>
		
		<button type="submit" class="btn btn-dark btn-sm">Submit</button>
	</form>`, Login: false, Payload: payload})
}

func Hard(w http.ResponseWriter, r *http.Request) {
	tmpl := htmlTemplate.Must(htmlTemplate.ParseFiles("templates/xss/xss_hard.html", "templates/base.html"))
	payload := r.FormValue("payload")

	tmpl.ExecuteTemplate(w, "xss_hard.html", struct {
		Payload string
		Login   bool
	}{Payload: payload, Login: false})

}
