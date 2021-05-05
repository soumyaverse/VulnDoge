package handler

import (
	"net/http"

	"github.com/burpOverflow/VulnDoge/api"
	"github.com/burpOverflow/VulnDoge/directoryTrversal"
	"github.com/burpOverflow/VulnDoge/loadimage"
)

func HandleDirectoryTrversal() {
	http.HandleFunc("/directoryTrversal/", directoryTrversal.DirectoryTrversalHandler)
	http.HandleFunc("/directoryTrversal/easy/", directoryTrversal.Easy)
	http.HandleFunc("/directoryTrversal/easy2/", directoryTrversal.Easy2)
	http.HandleFunc("/directoryTrversal/medium1/", directoryTrversal.Medium1)
	http.HandleFunc("/directoryTrversal/medium2/", directoryTrversal.Medium2)
	http.HandleFunc("/directoryTrversal/medium3/", directoryTrversal.Medium3)
}

func HandleAPI() {
	http.HandleFunc("/api/solution", api.APISolution)
}

func LoadImage() {
	http.HandleFunc("/loadimage", loadimage.LoadImageHandler)
	http.HandleFunc("/img", loadimage.Defult2)
}
