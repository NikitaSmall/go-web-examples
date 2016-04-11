package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nikitasmall/frameworks-comparsion/task"
)

var storage = task.NewStorage()

func initRouter(r *mux.Router) {
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/", addTaskHandler).Methods("POST")
	r.HandleFunc("/{id}", getTaskHandler).Methods("GET")
	r.HandleFunc("/{id}", deleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/{id}", updateTaskHandler).Methods("PUT")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tasks := storage.GetTasks()
	t, _ := json.Marshal(tasks)

	w.Write(t)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	storage.AddTask(text)

	http.Redirect(w, r, "/", 301)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	task, err := storage.GetTask(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		t, _ := json.Marshal(task)
		w.Write(t)
	}
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	task, err := storage.DeleteTask(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		t, _ := json.Marshal(task)
		w.Write(t)
	}
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	text := r.FormValue("text")
	complete, _ := strconv.ParseBool(r.FormValue("complete"))
	task, err := storage.UpdateTask(id, text, complete)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		t, _ := json.Marshal(task)
		w.Write(t)
	}
}

func main() {
	r := mux.NewRouter()
	initRouter(r)

	http.ListenAndServe(":3000", r)
}
