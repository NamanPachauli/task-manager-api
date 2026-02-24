package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Task struct {
	ID   int    `json:"id"`   // fixed tag
	Name string `json:"name"` // fixed tag
}

var tasks []Task // added (missing tha)

// home route
func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Task Manager API is running successfully") // fixed
}

// create task
func createTask(w http.ResponseWriter, r *http.Request) {

	var task Task
	taskhandle := json.NewDecoder(r.Body).Decode(&task)
	if taskhandle != nil {
		http.Error(w, taskhandle.Error(), http.StatusBadRequest) // fixed
		return
	}

	tasks = append(tasks, task) // moved inside function

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getTasks(w http.ResponseWriter, r *http.Request) { // fixed params

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks) // users → tasks
}

func updateTask(w http.ResponseWriter, r *http.Request) { // name fixed

	idParam := r.URL.Query().Get("id")
	id, Taskupdate := strconv.Atoi(idParam)

	if Taskupdate != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest) // fixed
		return
	}

	var updatedTask Task

	updatedhandle := json.NewDecoder(r.Body).Decode(&updatedTask)
	if updatedhandle != nil {
		http.Error(w, updatedhandle.Error(), http.StatusBadRequest) // fixed
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = updatedTask.Name // spelling fixed
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}

	http.Error(w, "user not found", http.StatusNotFound) // fixed
}

func deleteTask(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	id, deletehandle := strconv.Atoi(idParam)

	if deletehandle != nil { // spelling fixed
		http.Error(w, "invalid id", http.StatusBadRequest) // fixed
		return
	}

	for i, task := range tasks {
		if task.ID == id {

			tasks = append(tasks[:i], tasks[i+1:]...)

			w.WriteHeader(http.StatusOK) // fixed
			w.Write([]byte("User deleted successfully"))
			return
		}
	}

	http.Error(w, "user not found", http.StatusNotFound) // fixed
}

func main() {
	http.HandleFunc("/", server)
	http.HandleFunc("/create", createTask)
	http.HandleFunc("/tasks", getTasks)    // fixed name
	http.HandleFunc("/update", updateTask) // fixed name
	http.HandleFunc("/delete", deleteTask)

	fmt.Println("Server is running on port 9091...")
	http.ListenAndServe(":9091", nil)
}
