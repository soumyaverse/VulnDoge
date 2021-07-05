package csrf

import (
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

//CSRF where token is tied to non-session cookie
func Easy4(w http.ResponseWriter, r *http.Request) {
	// db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	// CheckErr.Check(err)
	// defer db.Close()

	// isSession, _ := SessionExist(r, db)
	// if isSession == true {
	// 	http.Redirect(w, r, "/csrf/easy3/myaccount/", 302)
	// 	return
	// }
	tmpl := template.Must(template.ParseFiles("templates/csrf/temp.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "temp.html", struct {
		Title string

		Login bool
		User  string
		Sol   bool
		Lid   string
	}{Title: "CSRF where token is not tied to user session", Login: false, Sol: true, Lid: "a5"})
}
