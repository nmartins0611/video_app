package root

import (
	"io/ioutil"
	"log"
	"net/http"

	"videoweb/config"

	"github.com/julienschmidt/httprouter"
)

//
// Index : Main landing page
func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	files, err := ioutil.ReadDir("./ui/static/videos")

	if err != nil {
		log.Fatal(err)
	}

	list := []string{}

	for _, f := range files {
		list = append(list, f.Name())
	}

	err = config.TPL.ExecuteTemplate(w, "index.ghtml", list)

	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
