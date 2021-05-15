package csrf

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/burpOverflow/VulnDoge/pkg/CheckErr"
	"github.com/burpOverflow/VulnDoge/pkg/rand"
	_ "github.com/go-sql-driver/mysql"
)

var csrfToken string

func Easy2(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	isSession, _ := SessionExist(r, db)
	if isSession == true {
		http.Redirect(w, r, "/csrf/easy2/myaccount/", 302)
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
	}{Title: "CSRF token validation depends on token being present", Desc: `<p style="color:green;">This Application validate the token when it is present but is skip validation if token is omitted. Try to hack change password functionality:) </p><div class="container"><h3>Create Account</h3>
	<form action='/csrf/easy2/create/' method='POST'>
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
	</form>or <a href='/csrf/easy2/login/'>Login</a>
	</div>`, Login: false, Sol: true, Lid: "a5"})
}

func CreateEasy2(w http.ResponseWriter, r *http.Request) {
	Create(w, r, "easy2")
}

func LoginEasy2(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	Login(w, r, "easy2", "CSRF token validation depends on token being present", db)
}

func MyAccountEasy2(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()
	isSession, uname := SessionExist(r, db)
	if isSession {
		csrfToken = rand.String(20)

		tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
		tmpl.ExecuteTemplate(w, "easy1.html", struct {
			Title     string
			Desc      string
			Login     bool
			User      string
			Sol       bool
			LogoutUrl string
			Lid       string
		}{Title: "CSRF token validation depends on token being present", Desc: `<h3>Welcome ` + uname + ` :)</h3><br><br><div class="container"><h4>Change Password</h4>
		<form action='/csrf/easy2/changepassword/' method='POST'>
		  <div class="mb-3">
		    <input type="hidden" name="csrf-token" value="` + csrfToken + `">
			<label for="newpassword" class="form-label">New Password</label>
			<input type="password" class="form-control" name="newpassword" required>
		  </div>
		  <button type="submit" class="btn btn-primary">Submit</button>
		</form>
		</div>`, Login: isSession, User: uname, LogoutUrl: "/csrf/easy2/logout/", Sol: false, Lid: "nil"})
	} else {
		http.Redirect(w, r, "/csrf/easy2/", 302)
	}
}

func LogoutEasy2(w http.ResponseWriter, r *http.Request) {
	Logout(w, r, "easy2")
}

func ChangePasswordEasy2(w http.ResponseWriter, r *http.Request) {
	newpassword := r.PostFormValue("newpassword")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	clientCsrfToken := r.PostFormValue("csrf-token")
	fmt.Println("server: ", csrfToken)
	fmt.Println("client: ", clientCsrfToken)
	fmt.Println()

	if clientCsrfToken == "" {
		fmt.Println("client csrf empty")
		_, uname := SessionExist(r, db)
		DBUpdatePassword(uname, newpassword, db)

		http.Redirect(w, r, "/csrf/easy2/", 302)
		return
	} else {
		if clientCsrfToken == csrfToken {
			fmt.Println("client csrf empty")
			_, uname := SessionExist(r, db)
			DBUpdatePassword(uname, newpassword, db)

			http.Redirect(w, r, "/csrf/easy2/", 302)
			return
		} else {
			fmt.Fprintf(w, "Invalid CSRF-TOKEN")
			return
		}
	}

	// _, uname := SessionExist(r, db)
	// DBUpdatePassword(uname, newpassword, db)

	// http.Redirect(w, r, "/csrf/easy2/", 302)

}
