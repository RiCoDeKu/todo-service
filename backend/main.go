package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Status	string	`json:"status"`
}

func main() {
	fmt.Println("Started Go Server")
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(todos); err != nil {
			http.Error(w, "failed to encode todos", http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}