package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VitalinaZhdanko/finalTask/app/controllers"
	"github.com/VitalinaZhdanko/finalTask/app/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	connect := database.InitDB()
	defer connect.Close()

	router := mux.NewRouter()

	router.HandleFunc("/groups", controllers.GetGroups).Methods(http.MethodGet)
	router.HandleFunc("/tasks", controllers.GetTasks).Methods(http.MethodGet)

	router.HandleFunc("/tasks/", controllers.PostTasks).Methods(http.MethodPost)
	router.HandleFunc("/groups/", controllers.PostGroups).Methods(http.MethodPost)
	router.HandleFunc("/timeframes/", controllers.PostTimeframes).Methods(http.MethodPost)

	router.HandleFunc("/tasks/{id}", controllers.PutTasks).Methods(http.MethodPut)
	router.HandleFunc("/groups/{id}", controllers.PutGroups).Methods(http.MethodPut)

	router.HandleFunc("/tasks/{id}", controllers.DeleteTasks).Methods(http.MethodDelete)
	router.HandleFunc("/groups/{id}", controllers.DeleteGroups).Methods(http.MethodDelete)
	router.HandleFunc("/timeframes/{id}", controllers.DeleteTimeframes).Methods(http.MethodDelete)

	fmt.Println("Server is listening...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
