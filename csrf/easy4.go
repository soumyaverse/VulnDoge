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

//CSRF where token is tied to non-session cookie
func Easy4(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	CheckErr.Check(err)
	defer db.Close()

	isSession, u := SessionExist(r, db)
	accountOption := ""
	if isSession == true {
		accountOption = `<p><a href="/csrf/easy4/myaccount/">My Account</a></p>`
	} else {
		accountOption = `<p><a href="/csrf/easy4/login/">Login</a> or <a href="/csrf/easy4/create/">Create Account</a></p>`
	}
	tmpl := template.Must(template.ParseFiles("templates/csrf/easy1.html", "templates/base.html"))

	if r.Method == http.MethodGet {

		search := r.FormValue("search")
		if search != "" {
			fmt.Println(search)

		}

		tmpl.ExecuteTemplate(w, "easy1.html", struct {
			Title string
			Desc  string
			Login bool
			User  string
			Sol   bool
			Lid   string
		}{Title: "CSRF where token is not tied to user session", Desc: accountOption + `
		<div class="container">
			<form action="/csrf/easy4/" method="get">
				<div class="input-group mb-3">
					<input type="text" class="form-control" placeholder="search here" aria-label="Recipient's username" aria-describedby="button-addon2" name="search" required>
					<button class="btn btn-primary" type="submit" id="button-addon2">Search</button>
				</div>
			</form>
		</div>`, Login: isSession, User: u, Sol: true, Lid: "a5"})
	}
}
