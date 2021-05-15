package csrf

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/burpOverflow/VulnDoge/pkg/CheckErr"
	_ "github.com/go-sql-driver/mysql"
)

func Pl() {
	fmt.Println("hello")
}

func Easy1(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	isSession, _ := SessionExist(r, db)
	if isSession == true {
		http.Redirect(w, r, "/csrf/easy1/myaccount/", 302)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "easy1.html", struct {
		Title string
		Desc  string
		Login bool
		User  string
		Sol   bool
		Lid   string
	}{Title: "csrf easy", Desc: `<p style="color:green;">It contains a very simple csrf vuln in change password functionality. Try to hack :) </p><div class="container"><h3>Create Account</h3>
	<form action='/csrf/easy1/create/' method='POST'>
	  <div class="mb-3">
		<div class="mb-3">
		<label for="username" class="form-label">Username</label>
		<input type="username" class="form-control" name="username" required>
	  </div>
		<label for="email" class="form-label">Email address</label>
		<input type="email" class="form-control" id="exampleInputEmail1" name="email" required>
		
	  </div>
	  <div class="mb-3">
		<label for="password" class="form-label">Password</label>
		<input type="password" class="form-control" name="password" required>
	  </div>
	  
	  <button type="submit" class="btn btn-primary">Submit</button>
	</form>or <a href='/csrf/easy1/login/'>Login</a>
	</div>`, Login: false, Sol: true, Lid: "a4"})
}

func LoginEasy1(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	Login(w, r, "easy1", "CSRF easy", db)
}

func CreateEasy1(w http.ResponseWriter, r *http.Request) {
	Create(w, r, "easy1")
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	newpassword := r.PostFormValue("newpassword")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()
	_, uname := SessionExist(r, db)
	DBUpdatePassword(uname, newpassword, db)

	http.Redirect(w, r, "/csrf/easy1/", 302)

}

func LogoutEasy1(w http.ResponseWriter, r *http.Request) {
	Logout(w, r, "easy1")
}
