package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/VitalinaZhdanko/finalTask/app/database"
	"github.com/VitalinaZhdanko/finalTask/app/models"
	"github.com/gorilla/mux"
)

//GetGroups handler for get groups
var GetGroups = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groups := database.ReadGroups()
	json.NewEncoder(w).Encode(groups)

}

// GetTasks handler
var GetTasks = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := database.ReadTasks()
	json.NewEncoder(w).Encode(tasks)
}

// PostTasks handler
var PostTasks = func(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	body, err := ioutil.ReadAll(r.Body)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&task)

	err = database.CreateTask(&task)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

// PostGroups handler
var PostGroups = func(w http.ResponseWriter, r *http.Request) {
	var group models.Group

	body, err := ioutil.ReadAll(r.Body)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&group)

	err = database.CreateGroup(&group)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

// PostTimeframes handler
var PostTimeframes = func(w http.ResponseWriter, r *http.Request) {

	var timeframe models.TimeFrame

	body, err := ioutil.ReadAll(r.Body)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&timeframe)

	err = database.CreateTimeframe(&timeframe)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(timeframe)
}

// PutTasks handler
var PutTasks = func(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	task.TaskID = id

	body, err := ioutil.ReadAll(r.Body)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&task)

	err = database.UpdateTask(&task)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// PutGroups handler
var PutGroups = func(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	group.GroupID = id

	body, err := ioutil.ReadAll(r.Body)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&group)

	err = database.UpdateGroup(&group)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}

// DeleteTasks handler
var DeleteTasks = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	err = database.DeleteTask(id)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusNoContent)

}

// DeleteGroups handler
var DeleteGroups = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	err = database.DeleteGroup(id)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteTimeframes handler
var DeleteTimeframes = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	err = database.DeleteTimeframes(id)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusNoContent)
}
