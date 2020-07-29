package main

import (
	"fmt"
	"net/http"

	"github.com/simonepm/dockerskin/hello"
)

const port = "8080"

func SayHelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", hello.SayHello())
}

func main() {

	http.HandleFunc("/", SayHelloServer)

	fmt.Println(fmt.Sprintf("Listening on port %s...", port))

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}
