package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hello World")
	})
        // Generate cert: go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
