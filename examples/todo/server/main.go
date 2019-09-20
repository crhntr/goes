package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/crhntr/goes/goesfixtures"
)

var (
	dependencies = os.Getenv("GOES_EXAMPLES_TODO_DEPENDENCIES_DIR")
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("listening on %q...", port)
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(Handler)))
}

func Handler(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL)

	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if strings.HasPrefix(req.URL.Path, "/dependencies/") {
		http.StripPrefix("/dependencies", http.FileServer(http.Dir(dependencies))).ServeHTTP(res, req)
		return
	}

	switch req.URL.Path {

	case goesfixtures.MinimalIndexPageExecutableWASMPath:
		res.Header().Set("content-type", "application/wasm")
		http.ServeFile(res, req, "webapp/main.wasm")

	case goesfixtures.MinimalIndexPageGoWASMExecPath:
		res.Header().Set("content-type", "application/json")
		res.Write(goesfixtures.GoWASMExec())

	default:
		res.WriteHeader(http.StatusOK)
		tmp, err := template.ParseFiles(filepath.Join(dependencies, "templates", "index.html"))
		if err != nil {
			log.Println(err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := tmp.Execute(res, struct {
			GoJS, WASMBin string
		}{
			goesfixtures.MinimalIndexPageGoWASMExecPath,
			goesfixtures.MinimalIndexPageExecutableWASMPath,
		}); err != nil {
			log.Println(err)
		}

	}
}
