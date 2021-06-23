package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/burpOverflow/VulnDoge/handler"
	"github.com/burpOverflow/VulnDoge/oAuth"
	"github.com/burpOverflow/VulnDoge/pkg/CheckErr"
	"github.com/burpOverflow/VulnDoge/xss"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConfig := flag.String("dbconfig", "nil", "mysql url")
	flag.Parse()
	if *dbConfig != "nil" {
		fmt.Println("setting database....")
		db, err := sql.Open("mysql", *dbConfig)
		CheckErr.Check(err)
		defer db.Close()

		_, err = db.Exec(`DROP DATABASE IF EXISTS VulnDoge`)
		CheckErr.Check(err)
		_, err = db.Exec(`CREATE DATABASE VulnDoge`)
		CheckErr.Check(err)
		db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
		CheckErr.Check(err)
		defer db.Close()
		_, err = db.Exec(`CREATE TABLE users(id INT PRIMARY KEY AUTO_INCREMENT,username VARCHAR(255),email VARCHAR(255),password VARCHAR(255),session VARCHAR(255),csrftoken VARCHAR(255))`)
		CheckErr.Check(err)
		_, err = db.Exec(`CREATE TABLE tokens(id INT PRIMARY KEY AUTO_INCREMENT,token VARCHAR(255))`)
		CheckErr.Check(err)

	}
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

	handler.HandleCSRF()

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
