package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type Book struct {
	id        string
	name      string
	completed bool
	due       time.Time
}

var db *sql.DB

func init() {
        var err error
	db, err = sql.Open("postgres", "user='' password='' dbname=todoapp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

        if err = db.Ping(); err != nil {
                log.Fatal(err)
        }
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
        router.POST("/todos", createBook)
        router.GET("/todos/:id", showBook)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
	})
	http.ListenAndServe(":8080", c.Handler(router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// db, err := sql.Open("postgres", "postgres://user:pass@localhost/todoapps")

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.id, &bk.name, &bk.completed, &bk.due)
		if err != nil {
                        http.Error(w, http.StatusText(500), 500)
                        return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
                http.Error(w, http.StatusText(500), 500)
                return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, %s\n", bk.id, bk.name, bk.completed, &bk.due)
	}
}

func showBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
        id := params.ByName("id")
        if id == "" {
                http.Error(w, http.StatusText(400), 400)
                return
        }
        row := db.QueryRow("SELECT * FROM todos WHERE id = $1", id)
        bk := new(Book)
        err := row.Scan(&bk.id, &bk.name, &bk.completed, &bk.due)
        if err != nil {
                http.Error(w, http.StatusText(500), 500)
                return
        }

        if err == sql.ErrNoRows {
                http.NotFound(w, r)
                return
        } else if err != nil {
                http.Error(w, http.StatusText(500), 500)
                return
        }
        fmt.Fprintf(w, "%s, %s, %s, %s\n", bk.id, bk.name, bk.completed, &bk.due)
}

func createBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        log.Printf("create Book")
        id := r.FormValue("id")
        name := r.FormValue("name")
        completed := r.FormValue("completed")
        due := r.FormValue("due")

        log.Printf(completed)
        if name == "" || completed == "" || due == "" {
                http.Error(w, http.StatusText(400), 400)
                return
        }

        result, err := db.Exec("INSERT INTO todos VALUES($1, $2, $3, $4)", id, name, completed, due)
        if err != nil {
                log.Print(err)
                http.Error(w, http.StatusText(500), 500)
                return
        }

        rowsAffected, err := result.RowsAffected()
        if err != nil {
                http.Error(w, http.StatusText(500), 500)
                return
        }
        fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", id, rowsAffected)
}
