package main

import (
	"fmt"
	"flag"
	"net/http"
	"log"
	"os/exec"
	"runtime"
)

type ConfigParams struct {
	port   *string
	folder *string
}

func main() {
	params := getConfigParams()

	go serveStaticContent(params)
	openRootFolderInBrowser(params)
	waitUntilKeyPress()
}
func getConfigParams() *ConfigParams {
	configParams := &ConfigParams{}

	configParams.port = flag.String("port", "8080", "The port in wich to run the static file web server")
	configParams.folder = flag.String("folder", "./www", "The name of the folder to be served")

	flag.Parse()
	return configParams
}

func serveStaticContent(params *ConfigParams) {
	http.Handle("/", http.FileServer(http.Dir(*params.folder)))

	if err := http.ListenAndServe(":" + *params.port, decorateDefaultHttpHandler()); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func openRootFolderInBrowser(params *ConfigParams) {
	openbrowser("http://localhost:" + *params.port)
}

func waitUntilKeyPress() {
	fmt.Println("Press <Enter> to stop serving files...")
	var end string
	fmt.Scanln(&end)
}

func decorateDefaultHttpHandler() http.Handler {
	return noCacheHandler(http.DefaultServeMux)
}

func noCacheHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("If-None-Match")
		r.Header.Del("If-Range")
		r.Header.Del("If-Modified-Since")

		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		if r.Proto == "HTTP/1.0" {
			w.Header().Set("Pragma", "no-cache")
		}

		handler.ServeHTTP(w, r)
	})
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
