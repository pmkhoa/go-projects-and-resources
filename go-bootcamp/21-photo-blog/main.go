package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("assets/templates/*.html")
}

type MyModel struct {
        IsLoggedin bool
        Images    []string
}

func getPhotos() []string {
	files := []string{}
	// or: make([]string, 0)
	filepath.Walk("./", func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}
		path = strings.Replace(path, "\\", "/", -1)
		if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".png") {
			files = append(files, path)
		}
		return nil
	})
	return files

}

func indexRoute(res http.ResponseWriter, req *http.Request) {
	// Parse template:
	session, _ := store.Get(req, "session")
        isLoggedin := false
        if session.Values["loggedin"] == true {
                isLoggedin = true 
        }
        err := tpl.ExecuteTemplate(res, "index.html", MyModel{
                IsLoggedin:      isLoggedin,
                Images:    getPhotos(),
        })
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func adminRoute(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if session.Values["loggedin"] == false || session.Values["loggedin"] == nil {
		http.Redirect(res, req, "/login", 302)
	} else {
		tpl.ExecuteTemplate(res, "admin.html", nil)
	}
	if req.Method == "POST" {
		src, hdr, err := req.FormFile("file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		// dst, err := os.Create(filepath.Join(os.TempDir(), hdr.Filename))
                wd, _ := os.Getwd()
                path := filepath.Join(wd, "assets", "images", hdr.Filename)
                dst, err := os.Create(path)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()
		io.Copy(dst, src)
	}
}

func logoutRoute(res http.ResponseWriter, req *http.Request) {
	// Remove login session
	session, _ := store.Get(req, "session")
	session.Values["loggedin"] = false
	session.Save(req, res)
	http.Redirect(res, req, "/", 302)
	return
}

func loginRoute(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if req.Method == "POST" && req.FormValue("password") == "secret" {
		session.Values["loggedin"] = true
		session.Save(req, res)
		http.Redirect(res, req, "/admin", 302)
		return
	}
	if session.Values["loggedin"] == true {
		http.Redirect(res, req, "/", 302)
	} else {
		tpl.ExecuteTemplate(res, "login.html", nil)
	}
}

func main() {
	http.HandleFunc("/", indexRoute)
	http.HandleFunc("/admin", adminRoute)
	http.HandleFunc("/logout", logoutRoute)
	http.HandleFunc("/login", loginRoute)
	http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("assets/images/"))))
	// http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
