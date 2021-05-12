package directoryTrversal

import (
	"net/http"
	"text/template"
)

func DirectoryTrversalHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/directoryTrversal/directoryTrversal.html", "templates/base.html"))

	tmpl.ExecuteTemplate(w, "directoryTrversal.html", nil)
}

func Easy(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/directoryTrversal/directoryTrversal_easy.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "directoryTrversal_easy.html", struct {
		Desc   string
		Title  string
		ImgSrc string
		Login  bool
	}{Desc: "Try to access <b>/etc/passwd</b> file using directoryTrversal vuln", Title: "Directory Trversal Easy", ImgSrc: "/loadimage?filename=go.png", Login: false})

}

func Easy2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/directoryTrversal/directoryTrversal_easy.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "directoryTrversal_easy.html", struct {
		Desc   string
		Title  string
		ImgSrc string
		Login  bool
	}{Desc: "Try to access <b>/etc/passwd</b> file using directoryTrversal Absolute Path vuln", Title: "Directory Trversal using Absolute path", ImgSrc: "/loadimage?filename=go.png", Login: false})

}

func Medium1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/directoryTrversal/directoryTrversal_easy.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "directoryTrversal_easy.html", struct {
		Login  bool
		Desc   string
		Title  string
		ImgSrc string
	}{Desc: "Try to access <b>/etc/passwd</b> file using directoryTrversal Where traversal sequences stripped non-recursively", Title: "Traversal sequences stripped", ImgSrc: "/loadimage?filename=go.png", Login: false})

}

func Medium2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/directoryTrversal/directoryTrversal_easy.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "directoryTrversal_easy.html", struct {
		Login  bool
		Desc   string
		Title  string
		ImgSrc string
	}{Desc: "Try to access <b>/etc/passwd</b> file using directoryTrversal Where traversal sequences are blocked and performs a URL-decode of the input before using it", Title: "Traversal sequences blocked with URL-decode", ImgSrc: "/loadimage?filename=go.png", Login: false})

}

func Medium3(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/directoryTrversal/directoryTrversal_easy.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "directoryTrversal_easy.html", struct {
		Login  bool
		Desc   string
		Title  string
		ImgSrc string
		Sol    bool
		Lid    string
	}{Desc: "Try to access <b>/etc/passwd</b> file using directoryTrversal Where The application transmits the full file path via a request parameter and validate whether it starts from <b>/var/www/images</b>", Title: "File path traversal, validation of start of path", ImgSrc: "/loadimage?filename=/var/www/images/go.png", Sol: true, Lid: "a1", Login: false})

}

func Medium4(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/directoryTrversal/directoryTrversal_easy.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "directoryTrversal_easy.html", struct {
		Path   string
		Desc   string
		Title  string
		ImgSrc string
		Sol    bool
		Lid    string
		Login  bool
	}{Path: "» Easy", Desc: "Try to access <b>/etc/passwd</b> file using directoryTrversal Where The application validates the file extensation", Title: "File path traversal, validation of file extensation", ImgSrc: "/loadimage?filename=go.png", Sol: true, Lid: "a2", Login: false})

}
