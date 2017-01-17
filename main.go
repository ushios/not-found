package main

import (
	"flag"
	"fmt"
	"net/http"
)

const DefaultPort = 80

var port int

func main() {
	flag.IntVar(&port, "port", DefaultPort, "http server port")
	flag.IntVar(&port, "p", DefaultPort, "http server port")
	flag.Parse()

	fmt.Println("listen port", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
