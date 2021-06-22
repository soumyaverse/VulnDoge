package csrf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"net/http"
	"os"
	"text/template"

	"github.com/burpOverflow/VulnDoge/pkg/CheckErr"
	"github.com/burpOverflow/VulnDoge/pkg/rand"
)

func Easy3(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	isSession, _ := SessionExist(r, db)
	if isSession == true {
		http.Redirect(w, r, "/csrf/easy3/myaccount/", 302)
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
	}{Title: "CSRF where token is not tied to user session", Desc: `<p style="color:green;">Here CSRF where token is not tied to user session. Try to find out CSRF on change password functionality :)</p><div class="container"><h3>Create Account</h3>
	<form action='/csrf/easy3/create/' method='POST'>
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
	</form>or <a href='/csrf/easy3/login/'>Login</a>
	</div>`, Login: false, Sol: true, Lid: "a5"})
}

func CreateEasy3(w http.ResponseWriter, r *http.Request) {
	Create(w, r, "easy3")
}

func LoginEasy3(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	Login(w, r, "easy3", "CSRF where token is not tied to user session", db)
}

func MyAccountEasy3(w http.ResponseWriter, r *http.Request) {

	token := rand.String(16)
	TokenInsert(token)
	MyAccountCSRFToken(w, r, "CSRF where token is not tied to user session", "easy3", token)

}

func LogoutEasy3(w http.ResponseWriter, r *http.Request) {
	Logout(w, r, "easy3")
}
