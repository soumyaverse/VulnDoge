package oAuth

import (
	"fmt"
	"net/http"
)

func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello OAuth!")
}
