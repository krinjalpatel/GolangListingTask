package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The Task Type (more like an object)
type Task struct {
	ID   string `json:"id,omitempty"`
	Task string `json:"task,omitempty"`
	Time string `json:"time,omitempty"`
}

var task []Task

// GetList gets all the tasks in the list
func GetList(w http.ResponseWriter, r *http.Request) {
	userTask, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userTask)
}

// GetTask displays a particular the task in the list
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range task {
		if item.ID == params["id"] {
			userTask, err := json.Marshal(item)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(userTask)
			return
		}
	}
}

// CreateTask creates a new task in the list
func CreateTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newTask := Task{ID: r.FormValue("id"), Task: r.FormValue("task"), Time: r.FormValue("time")}
	task = append(task, newTask)
	userTask, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userTask)
}

// DeleteTask deletes an item from the list
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range task {
		if string(item.ID) == string(params["id"]) {
			task = append(task[:index], task[index+1:]...)
			break
		}
	}
	userTask, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userTask)
}

func main() {
	router := mux.NewRouter()
	task = append(task, Task{ID: "1", Task: "To iron clothes", Time: "6:00 pm"})
	task = append(task, Task{ID: "2", Task: "To create science project of niece", Time: "7:00 pm"})
	task = append(task, Task{ID: "3", Task: "To chop vegetables", Time: "9:00 pm"})
	task = append(task, Task{ID: "4", Task: "To pay Electricity bill online", Time: "9:30 pm"})
	task = append(task, Task{ID: "5", Task: "To transfer money to Alia", Time: "9:45 pm"})
	router.HandleFunc("/getList", GetList).Methods("GET")
	router.HandleFunc("/getTask/{id}", GetTask).Methods("GET")
	router.HandleFunc("/createTask", CreateTask).Methods("POST")
	router.HandleFunc("/deleteTask/{id}", DeleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8010", router))
}
