package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("port", "1088", "server port")

func main() {
	flag.Parse()
	fmt.Println("Start server with docs on port ", *port)
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("./docs"))))
}
