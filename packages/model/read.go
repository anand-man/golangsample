package model

import "github.com/anand-man/golangsample/packages/views"

func ReadAll() ([]views.PostRequest, error) {
	row, err := con.Query("SELECT * FROM TODO")

	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}

	for row.Next() {
		data := views.PostRequest{}
		row.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	row, err := con.Query("SELECT * FROM TODO WHERE name=?", name)

	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}

	for row.Next() {
		data := views.PostRequest{}
		row.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}
