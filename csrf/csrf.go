package csrf

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/burpOverflow/VulnDoge/pkg/CheckErr"
	"github.com/burpOverflow/VulnDoge/pkg/rand"
	_ "github.com/go-sql-driver/mysql"
)

func CSRFHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/csrf/csrf.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "csrf.html", nil)

}

func Login(w http.ResponseWriter, r *http.Request, base string, title string, db *sql.DB) {

	if r.Method == http.MethodGet {

		isSession, uname := SessionExist(r, db)
		if isSession == true {
			http.Redirect(w, r, "/csrf/"+base+"/", 302)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
		tmpl.ExecuteTemplate(w, "easy1.html", struct {
			Title string
			Desc  string
			Login bool
			User  string
		}{Title: title, Desc: `<div class="container"><h3>Login</h3>
		<form action='/csrf/` + base + `/login/' method='POST'>
		<div class="mb-3">
			<div class="mb-3">
			<label for="username" class="form-label">Username</label>
			<input type="username" class="form-control" name="username" required>
		  </div>			
		</div>
		  <div class="mb-3">
			<label for="password" class="form-label">Password</label>
			<input type="password" class="form-control" name="password" required>
		  </div>
		  <button type="submit" class="btn btn-primary">Submit</button>
		</form>or <a href='/csrf/` + base + `/'>Create Account</a>
		</div>`, Login: isSession, User: uname})
	}
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		var dbpassword string
		err2 := db.QueryRow(`SELECT password FROM users WHERE username = ? `, username).Scan(&dbpassword)
		CheckErr.Check(err2)

		if password == dbpassword {
			StoreCookie(w, db, username, base)
			http.Redirect(w, r, "/csrf/"+base+"/login/", 302)
		}

		fmt.Fprintf(w, "not logged in!")
	}
}

func Create(w http.ResponseWriter, r *http.Request, base string) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	if r.Method == http.MethodGet {
		isSession, _ := SessionExist(r, db)
		if isSession == true {
			http.Redirect(w, r, "/csrf/"+base+"/myaccount/", 302)
		}
	}
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		check := UserExists(db, username)
		if check == true {
			fmt.Fprintf(w, username+" user exists")
			return
		}
		check = EmailExists(db, email)
		if check == true {
			fmt.Fprintf(w, email+" email exists")
			return
		}

		sql := `INSERT INTO users(username,email,password) VALUES(?,?,?)`
		_, err = db.Exec(sql, username, email, password)
		CheckErr.Check(err)
		StoreCookie(w, db, username, base)

		http.Redirect(w, r, "/csrf/"+base+"/", 302)
	}
}

func StoreCookie(w http.ResponseWriter, db *sql.DB, username string, base string) {
	// cookie, err := r.Cookie("session")
	// fmt.Println("cookie: ", cookie)
	// if err != nil {
	// 	fmt.Println("cookie was not found")
	// 	cookieValue := rand.String(16)
	// 	cookie = &http.Cookie{
	// 		Name:     "session",
	// 		Value:    cookieValue,
	// 		HttpOnly: false,
	// 		Path:     "/csrf/",
	// 	}
	// 	http.SetCookie(w, cookie)

	// users.Users[i].Cookie = cookieValue
	// fmt.Println(users.Users[i].Cookie)
	// fmt.Println(users)
	// DbUpdate(users)
	// }
	cookieValue := rand.String(16)
	cookie := &http.Cookie{
		Name:     "session",
		Value:    cookieValue,
		HttpOnly: false,
		Path:     "/csrf/" + base + "/",
	}
	DBUpdateSession(username, cookieValue, db)
	http.SetCookie(w, cookie)

}

func UserExists(db *sql.DB, username string) bool {
	sqlStmt := `SELECT username FROM users WHERE username = ?`
	err := db.QueryRow(sqlStmt, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {

			log.Print(err)
		}

		return false
	}

	return true
}

func EmailExists(db *sql.DB, email string) bool {
	sqlStmt := `SELECT email FROM users WHERE email = ?`
	err := db.QueryRow(sqlStmt, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {

			log.Print(err)
		}

		return false
	}

	return true
}

func TokenExists(db *sql.DB, token string) bool {
	sqlquery := `SELECT token FROM tokens WHERE token = ?`
	err := db.QueryRow(sqlquery, token).Scan(&token)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

func TokenInsert(token string) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)

	_, err = db.Exec(`INSERT INTO tokens(token) VALUE(?)`, token)
	CheckErr.Check(err)
	defer db.Close()
}

func DBUpdateSession(username string, session string, db *sql.DB) {
	_, err := db.Exec(`UPDATE users SET session= ? WHERE username= ? `, session, username)
	CheckErr.Check(err)
	defer db.Close()

}

func DBUpdatePassword(username string, newpass string, db *sql.DB) {
	_, err := db.Exec(`UPDATE users SET password= ? WHERE username= ? `, newpass, username)
	CheckErr.Check(err)
	// defer db.Close()
}

func SessionExist(r *http.Request, db *sql.DB) (bool, string) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false, "nil"
	}
	var uname string
	err = db.QueryRow(`SELECT username FROM users WHERE session = ? `, cookie.Value).Scan(&uname)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println(err)
		}
		return false, "nil"
	}
	if len(strings.Trim(uname, "")) > 1 && len(cookie.Value) == 16 {
		fmt.Println("Session Exist: " + uname)
		return true, uname
	}

	return false, "nil"
}

func MyAccount(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()
	isSession, uname := SessionExist(r, db)
	if isSession {

		tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
		tmpl.ExecuteTemplate(w, "easy1.html", struct {
			Title     string
			Desc      string
			Login     bool
			User      string
			LogoutUrl string
		}{Title: "csrf easy", Desc: `<h3>Welcome ` + uname + ` :)</h3><br><br><div class="container"><h4>Change Password</h4>
		<form action='/csrf/easy1/changepassword/' method='POST'>
		  <div class="mb-3">
			<label for="newpassword" class="form-label">New Password</label>
			<input type="password" class="form-control" name="newpassword" required>
		  </div>
		  <button type="submit" class="btn btn-primary">Submit</button>
		</form>
		</div>`, Login: isSession, User: uname, LogoutUrl: "/csrf/easy1/logout/"})
	} else {
		http.Redirect(w, r, "/csrf/easy1/", 302)
	}
}

func MyAccountCSRFToken(w http.ResponseWriter, r *http.Request, title string, base string, token string) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)

	isSession, uname := SessionExist(r, db)
	fmt.Println("MyAccountCSRFToken: ", uname)
	fmt.Println("isSession: ", isSession)

	if isSession {

		tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
		tmpl.ExecuteTemplate(w, "easy1.html", struct {
			Title     string
			Desc      string
			Login     bool
			User      string
			Sol       bool
			LogoutUrl string
			Lid       string
		}{Title: title, Desc: `<h3>Welcome ` + uname + ` :)</h3><br><br><div class="container"><h4>Change Password</h4>
		<form action='/csrf/` + base + `/changepassword/' method='POST'>
		  <div class="mb-3">
		    <input type="hidden" name="csrf-token" value="` + token + `">
			<label for="newpassword" class="form-label">New Password</label>
			<input type="password" class="form-control" name="newpassword" required>
		  </div>
		  <button type="submit" class="btn btn-primary">Submit</button>
		</form>
		</div>`, Login: isSession, User: uname, LogoutUrl: "/csrf/" + base + "/logout/", Sol: false, Lid: "nil"})
	} else {
		http.Redirect(w, r, "/csrf/"+base+"/", 302)
	}
}

func Logout(w http.ResponseWriter, r *http.Request, base string) {
	clearSession(w, "/csrf/"+base+"/")
	http.Redirect(w, r, "/csrf/"+base+"/", 302)
}

func clearSession(w http.ResponseWriter, path string) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   path,
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
