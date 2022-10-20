package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anand-man/golangsample/packages/model"

	"github.com/anand-man/golangsample/packages/controller"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Server start")
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}

// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	fmt.Println("Go MySQL Tutorial")

// 	// Open up our database connection.
// 	// I've set up a database on my local machine using phpmyadmin.
// 	// The database is called testDb
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")

// 	// if there is an error opening the connection, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Start")
// 	// defer the close till after the main function has finished
// 	// executing
// 	defer db.Close()
// 	fmt.Println("Stop")

// }
