package csrf

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/burpOverflow/VulnDoge/pkg/CheckErr"
)

func Pl() {
	fmt.Println("hello")
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
