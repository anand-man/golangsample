package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anand-man/golangsample/packages/model"

	"github.com/anand-man/golangsample/packages/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error!"))
				return
			} else {
				fmt.Println("data created!")
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			data, err := model.ReadAll()
			fmt.Println(r.URL.Query())
			if len(r.URL.Query()) == 1 {
				data, err = model.ReadByName(name)
			}
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]

			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"Item deleted!"})
		}

	}
}
