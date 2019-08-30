// +build acceptance

package greeting_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/crhntr/goes/goestest"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("go", strings.Fields("build -o webapp/main.wasm webapp/main.go")...)
	cmd.Env = append(os.Environ(), strings.Fields("GOOS=js GOARCH=wasm")...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestHelloBox_Acceptance(t *testing.T) {
	respondToSpanishGreetingRequest, respondedToSpanishGreetingRequest := make(chan struct{}), make(chan struct{})

	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
			defer close(respondedToSpanishGreetingRequest)
			<-respondToSpanishGreetingRequest

			res.Header().Set("content-type", "text/plain")
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("¡Hola, mundo!"))
		case "/":
			res.WriteHeader(http.StatusOK)
			res.Write(goestest.MinimalIndexPage())
		default:
			res.WriteHeader(http.StatusNotFound)
		}
	}))

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *runtime.EventConsoleAPICalled:
			fmt.Printf("console.%s call: ", ev.Type)
			for _, arg := range ev.Args {
				fmt.Printf("%s (%s), ", arg.Value, arg.Type)
			}
			fmt.Println()
		}
	})

	var (
		message, spanishMessage string
	)
	err := chromedp.Run(ctx,
		chromedp.Navigate(server.URL),
		chromedp.Text(`hello-box`, &message, chromedp.NodeVisible, chromedp.ByID),
		chromedp.ActionFunc(func(context.Context) error {
			close(respondToSpanishGreetingRequest)
			<-respondedToSpanishGreetingRequest
			return nil
		}),
		chromedp.Text(`hello-box`, &spanishMessage, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	if message != "Hello, world!" {
		t.Error("it should set the inital greeting")
		t.Logf("got: %q", message)
	}
	if !strings.Contains(spanishMessage, "Hola") {
		t.Error("it should set the server's response greeting")
		t.Logf("got: %q", spanishMessage)
	}
}
