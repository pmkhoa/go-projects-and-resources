package todo

import "net/http"

// func indexRoute(res http.ResponseWriter, req *http.Request) {
// 	tpl, err := template.ParseFiles("assets/templates/index.gohtml")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	err = tpl.ExecuteTemplate(res, "index.gohtml", nil)
// 	if err != nil {
// 		http.Error(res, err.Error(), 500)
// 	}
// }

func indexRoute(res http.ResponseWriter, req *http.Request) {
        if req.URL.Path != "/" {
                http.NotFound(res, req)
                return
        }
	http.ServeFile(res, req, "assets/templates/index.html")
}

func handleTodos(res http.ResponseWriter, req *http.Request) {
        switch req.Method {
                case "GET":

                case "POST":

                case "DELETE":

                default:
                        http.Error(res, "Method not allow", 405)
        }
}

func init() {
        http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", indexRoute)
        http.HandleFunc("/todo.json", handleTodos)
}
