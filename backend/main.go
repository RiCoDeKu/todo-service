package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Status	string	`json:"status"`
}

type CreateTodoRequest struct {
	Title	string	`json:"title"`
	Status	string	`json:"status"`
}

func main() {
	fmt.Println("Started Go Server...")
	todos := []Todo{
	{
		ID:	1,
		Title:  "Reactを勉強する",
		Status: "done",
	},
	{
		ID:     2,
		Title:  "Goを勉強する",
		Status: "todo",
	},
}

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		// GET Process
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")

			if err := json.NewEncoder(w).Encode(todos); err != nil {
				http.Error(w, "failed to encode todos", http.StatusInternalServerError)
				return
			}

		// POST Process
		case http.MethodPost:
			var req CreateTodoRequest

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			newTodo := Todo{
				ID:		len(todos) + 1,
				Title:	req.Title,
				Status: req.Status,
			}

			todos = append(todos, newTodo)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			if err := json.NewEncoder(w).Encode(newTodo); err != nil {
				return
			}

		// Other Process
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET: fetch specific todo.
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet{
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idText := strings.TrimPrefix(r.URL.Path, "/todos/")

		id, err := strconv.Atoi(idText)
		if err != nil {
			http.Error(w, "invalid todo id", http.StatusBadRequest)
			return
		}

		for _, todo := range todos {
			if todo.ID == id {
				w.Header().Set("Content-Type", "application/json")

				if err := json.NewEncoder(w).Encode(todo); err != nil {
					http.Error(w, "failed to encode todo", http.StatusInternalServerError)
				}
				return
			}
		}

		http.Error(w, "todo not found", http.StatusInternalServerError)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}