package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", fakeWebServer)
	http.ListenAndServe(":8000", nil)
}

func fakeWebServer(w http.ResponseWriter, r *http.Request) {
	dummy()
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func dummy() {
	x := 0.0001
	for i := 0; i <= 10000000; i++ {
		x += math.Sqrt(x)
	}
}

func greeting(msg string) string {
	return ("<b>" + msg + "</b>")
}
