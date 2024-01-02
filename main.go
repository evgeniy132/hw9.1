package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Task struct {
	Date        string `json:"date"`
	Description string `json:"description"`
}

var tasks = []Task{
	{Date: "2024-01-01", Description: "Вправа 1"},
	{Date: "2024-01-01", Description: "Вправа 2"},
	{Date: "2024-01-02", Description: "Вправа 3"},
	// Додайте тут інші завдання за потреби
}

var mutex sync.RWMutex

func main() {
	http.HandleFunc("/tasks", handleTasks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
		return
	}

	date := r.URL.Query().Get("date")

	mutex.RLock()
	defer mutex.RUnlock()

	filteredTasks := []Task{}
	for _, task := range tasks {
		if task.Date == date {
			filteredTasks = append(filteredTasks, task)
		}
	}

	jsonTasks, err := json.Marshal(filteredTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonTasks)
}
