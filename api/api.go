package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Solutions struct {
	Solutions []Solution `json:"solutions"`
}

type Solution struct {
	Lid   string `json:"lid"`
	Title string `json:"title"`
	Sol   string `json:"sol"`
}

func APISolution(w http.ResponseWriter, r *http.Request) {
	lid := r.FormValue("lid")
	fmt.Println(lid)

	jsonfile, err := os.Open("db/solutions.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonfile.Close()
	var solutions Solutions
	jsondata, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(jsondata, &solutions)
	fmt.Println("Solution: ", r.Referer())
	w.Header().Add("Content-Type", "text")

	for i := 0; i < len(solutions.Solutions); i++ {
		if solutions.Solutions[i].Lid == lid {
			fmt.Fprintf(w, solutions.Solutions[i].Sol)
			return
		}
	}
	fmt.Fprintf(w, "Unable to find solution")
}
