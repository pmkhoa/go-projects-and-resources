package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func TodosIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	id, err := strconv.Atoi(todoId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	todo := FindTodo(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		panic(err)
	}
}

// curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &todo)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			panic(err)
		}
	}
	t := CreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(t)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// curl -X DELETE 'http://localhost:8080/todos/2'
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	id, err := strconv.Atoi(todoId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
        err = DestroyTodo(id)
        if err != nil {
		http.Error(w, err.Error(), 500)
		return
        }
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
