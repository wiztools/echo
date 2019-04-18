package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var str strings.Builder

		// Request:
		var b bytes.Buffer
		wrtr := bufio.NewWriter(&b)
		r.Write(wrtr)
		wrtr.Flush()
		str.WriteString(string(b.Bytes()))

		fmt.Fprintf(w, str.String())
	})
	http.ListenAndServe(":3000", nil)
}
