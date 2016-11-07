package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
                if req.Method == "POST" {
                        key := "q"
                        // val := req.URL.Query().Get(key)
                        // io.WriteString(res, "Value: "+val)
                        // formValues := req.FormValue(key)
                        file, _, err := req.FormFile(key)
                        if err != nil {
                                http.Error(res, err.Error(), 500)
                                return
                        }
                        defer file.Close()

                        bs, err := ioutil.ReadAll(file)
                        fmt.Println(string(bs))
                }
		// fmt.Println("Value: ", formValues)
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `<form method="POST" enctype="multipart/form-data">
                       <input type="file" name="q">
                       <input type="submit" value="submit">
                       </form>`)
	})
	http.ListenAndServe(":9000", nil)
}

