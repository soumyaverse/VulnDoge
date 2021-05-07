package csrf

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/burpOverflow/VulnDoge/pkg/rand"
)

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
		fmt.Println(username)
		fmt.Println(password)
		fmt.Fprintf(w, "success login")
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(password)
		fmt.Fprintf(w, "success create")
	}
}

func StoreCookie(w http.ResponseWriter, r *http.Request, i int) {
	cookie, err := r.Cookie("session")
	fmt.Println("cookie: ", cookie)
	if err != nil {
		fmt.Println("cookie was not found")
		cookieValue := rand.String(16)
		cookie = &http.Cookie{
			Name:     "session",
			Value:    cookieValue,
			HttpOnly: false,
			Path:     "/csrf/",
		}
		http.SetCookie(w, cookie)

		// users.Users[i].Cookie = cookieValue
		// fmt.Println(users.Users[i].Cookie)
		// fmt.Println(users)
		// DbUpdate(users)
	}

}
