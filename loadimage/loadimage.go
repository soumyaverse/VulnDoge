package loadimage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func LoadImageHandler(w http.ResponseWriter, r *http.Request) {
	imagename := r.FormValue("filename")
	PORT := os.Getenv("VPORT")

	addr := []string{"http://localhost:" + PORT + "/directoryTrversal/", "http://127.0.0.1:" + PORT + "/directoryTrversal/"}

	if r.Referer() == addr[0]+"easy/" || r.Referer() == addr[1]+"easy/" {

		if imagename == "../../../etc/passwd" {
			file, _ := ioutil.ReadFile("media/passwd.txt")
			fmt.Fprintf(w, string(file))
			return
		}

		defult(imagename, w)
	}
	if r.Referer() == addr[0]+"easy2/" || r.Referer() == addr[1]+"easy2/" {

		if imagename == "/etc/passwd" {
			file, _ := ioutil.ReadFile("media/passwd.txt")
			fmt.Fprintf(w, string(file))
			return
		}

		defult(imagename, w)
	}
	if r.Referer() == addr[0]+"medium1/" || r.Referer() == addr[1]+"medium1/" {

		if imagename == "....//....//....//etc/passwd" {
			file, _ := ioutil.ReadFile("media/passwd.txt")
			fmt.Fprintf(w, string(file))
			return
		}

		defult(imagename, w)
	}
	if r.Referer() == addr[0]+"medium2/" || r.Referer() == addr[1]+"medium2/" {
		fmt.Println("medium2 img load..")
		if imagename == "go.png" {
			defult(imagename, w)
		} else {
			fmt.Println("img name match")
			file, _ := ioutil.ReadFile("media/passwd.txt")
			fmt.Fprintf(w, string(file))
			return
		}

		return
	}

	if r.Referer() == addr[0]+"medium3/" || r.Referer() == addr[1]+"medium3/" {
		if imagename[0:16] == "/var/www/images/" {

			if imagename == "/var/www/images/../../../etc/passwd" {
				file, _ := ioutil.ReadFile("media/passwd.txt")
				fmt.Fprintf(w, string(file))
				return
			}

			defult(imagename[16:], w)
		}
		fmt.Fprint(w, "Invalid Path")
		return

	}
	defult(imagename, w)

}

func checkImgReadErr(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Fprintf(w, "Image Not Found")
		fmt.Println(err)
		return
	}
}

func defult(imagename string, w http.ResponseWriter) {
	img, err := ioutil.ReadFile("media/" + imagename)
	checkImgReadErr(err, w)

	fmt.Fprint(w, string(img))
	return

}
