package csrf

import (
	"fmt"
	"net/http"

	"github.com/burpOverflow/VulnDoge/pkg/rand"
)

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
