package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Request:
		var b bytes.Buffer
		wrtr := bufio.NewWriter(&b)
		r.Write(wrtr)
		wrtr.Flush()

		fmt.Fprintf(w, string(b.Bytes()))
	})
	http.ListenAndServe(":3000", nil)
}
