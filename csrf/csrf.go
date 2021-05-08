package csrf

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/burpOverflow/VulnDoge/pkg/rand"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id       int
	Username string
	Email    string
	Password string
}

func CSRFHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/csrf/csrf.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "csrf.html", nil)

}

func Easy1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "easy1.html", struct {
		Title string
		Desc  string
	}{Title: "csrf easy", Desc: "<h3>Create Account</h3><form action='/csrf/easy1/create/' method='POST'><label for='username'>Username: </label><input type='text' name='username'><br><label for='email'>Email: &nbsp;&nbsp;&nbsp;</label><input type='email' name='email'><br><label for='password'>Password: </label><input type='password' name='password'><br><br><button type='submit'>Create</button></form> or <a href='/csrf/easy1/login/'>Login</a>"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))
		tmpl.ExecuteTemplate(w, "easy1.html", struct {
			Title string
			Desc  string
		}{Title: "csrf easy", Desc: "<h3>Login</h3><form action='/csrf/easy1/login/' method='POST'><label for='username'>Username: </label><input type='text' name='username'><br><label for='password'>Password: </label><input type='password' name='password'><br><br><button type='submit'>Login</button></form> or <a href='/csrf/easy1/'>Create Account</a>"})
	}
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		var dbpassword string
		err2 := db.QueryRow(`SELECT password from users WHERE username = ? `, username).Scan(&dbpassword)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println("password: ", password)
		fmt.Println("dbpassword: ", dbpassword)
		if password == dbpassword {
			StoreCookie(w)
			http.Redirect(w, r, "/csrf/easy1/", 302)
		}
		// sql := `SELECT * FROM users WHERE username = ? `
		// res, err := db.Query(sql, username)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer res.Close()
		// if res.Next() {
		// 	var users Users
		// 	err := res.Scan(&users.Id, &users.Username, &users.Email, &users.Password)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println(users)
		// }

		fmt.Fprintf(w, "not logged in!")
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

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

		// sql := "INSERT INTO users(username,email,password) VALUES(" + "'" + username + "'" + "," + "'" + email + "'" + "," + "'" + password + "'" + ")"

		sql := `INSERT INTO users(username,email,password) VALUES(?,?,?)`
		_, err = db.Exec(sql, username, email, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(password)
		fmt.Fprintf(w, "success create")
	}
}

func StoreCookie(w http.ResponseWriter) {
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
		Path:     "/csrf/easy1/",
	}
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
