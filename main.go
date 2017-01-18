package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

const (
	DefaultPort            = 80
	DefaultHelathCheckPath = "/healthcheck"
)

var (
	port        int
	healthcheck string
)

func main() {
	flag.IntVar(&port, "port", DefaultPort, "http server port")
	flag.IntVar(&port, "p", DefaultPort, "http server port")
	flag.StringVar(&healthcheck, "healthcheck", DefaultHelathCheckPath, "health check path (return http status code 200)")
	flag.StringVar(&healthcheck, "c", DefaultHelathCheckPath, "health check path (return http status code 200)")
	flag.Parse()

	healthcheck = filepath.Join("/", healthcheck)

	fmt.Println("listen port:       ", port)
	fmt.Println("health check path: ", healthcheck)

	http.HandleFunc(healthcheck, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
