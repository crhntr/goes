package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/crhntr/goes"
	"github.com/crhntr/goes/examples/greeting"
)

func main() {
	port := os.Getenv("8080")
	if port == "" {
		port = "8080"
	}

	assets := os.Getenv("ASSETS_DIRECTORY")
	if assets == "" {
		assets = "assets"
	}

	fileServer := http.FileServer(http.Dir(assets))

	indexPath := filepath.Join(assets, "index.gohtml")

	log.Printf("listening on %q...", port)

	log.Fatal(http.ListenAndServe(":"+port, Handler(indexPath, fileServer)))
}

func Handler(indexPath string, fileServer http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		switch req.URL.Path {
		case "/api/spanish-greeting":
			time.Sleep(time.Second * 2)

			res.Header().Set("content-type", "text/plain")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("¡Hola, mundo!"))

		case "/":
			f, err := os.Open(indexPath)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
			defer f.Close()

			tmpl, err := goes.New(path.Base(indexPath), f)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}

			data := struct {
				greeting.Envelope
				*goes.Template
			}{
				Envelope: greeting.Envelope{Message: "Hello, world!"},
				Template: tmpl,
			}

			if err := tmpl.ExecuteTemplate(res, path.Base(indexPath), data); err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}

		default:
			if path.Ext(req.URL.Path) == "wasm" {
				res.Header().Set("content-type", "application/wasm")
			}

			fileServer.ServeHTTP(res, req)
		}
	}
}
