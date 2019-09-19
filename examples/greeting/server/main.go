package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/crhntr/goes/goestest"
)

func main() {
	listen := flag.String("listen", ":8080", "listen address")
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.HandlerFunc(Handler)))
}

func Handler(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL)
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	switch req.URL.Path {
	case goestest.MinimalIndexPageExecutableWASMPath:
		res.Header().Set("content-type", "application/wasm")
		http.ServeFile(res, req, "webapp/main.wasm")
	case goestest.MinimalIndexPageGoWASMExecPath:
		res.Header().Set("content-type", "application/json")
		res.Write(goestest.GoWASMExec())
	case "/api/spanish-greeting":
		res.Header().Set("content-type", "text/plain")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("¡Hola, mundo!"))
	case "/":
		res.WriteHeader(http.StatusOK)
		res.Write(goestest.MinimalIndexPage())
	default:
		res.WriteHeader(http.StatusNotFound)
	}
}