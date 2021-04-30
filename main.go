package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/burpOverflow/VulnDoge/handler"
	"github.com/burpOverflow/VulnDoge/oAuth"
	"github.com/burpOverflow/VulnDoge/pkg/CheckErr"
	"github.com/burpOverflow/VulnDoge/xss"
)

func main() {
	PORT := os.Getenv("VPORT")

	http.HandleFunc("/", Index)
	http.HandleFunc("/xss/", xss.XSSHandler)
	http.HandleFunc("/xss/easy/", xss.Easy)
	http.HandleFunc("/xss/hard/", xss.Hard)

	http.HandleFunc("/oauth/", oAuth.OAuthHandler)

	// http.HandleFunc("/csrf/", csrf.CSRFHandler)
	// http.HandleFunc("/csrf/create/", csrf.Create)
	// http.HandleFunc("/csrf/login/", csrf.Login)

	// for directoryTrversal
	handler.HandleDirectoryTrversal()

	handler.LoadImage()

	handler.HandleAPI()

	fmt.Println(string("\033[36m"), "[+] Started on http://127.0.0.1:"+PORT, string("\033[0m"))
	err := http.ListenAndServe(":"+PORT, nil)
	CheckErr.Check(err)

}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	CheckErr.Check(err)
}
