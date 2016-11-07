package main

import (
	"io"
	"net/http"
	"os"
)

func helloWorld(res http.ResponseWriter, req *http.Request) {
	requestType := req.Method
	io.WriteString(res, `
                Hello World
                RequestType: `+requestType+`
        `)
}

func gifLocalCat(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("giphy.gif")
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}

func loadLocalGifCat(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "giphy.gif")
}

func gifCat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `
       <!DOCTYPE html>
        <head></head>
        <body>
                <img src="/giphy.gif">
        </body>

       </html>`)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/cat", gifCat)
	http.HandleFunc("/giphy.gif", loadLocalGifCat)
	http.ListenAndServe(":9000", nil)
}
