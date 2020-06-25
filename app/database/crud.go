package database

import (
	"database/sql"
	"os"

	"github.com/VitalinaZhdanko/finalTask/app/models"
)

type dbW struct {
	db *sql.DB
}

var repo = dbW{}

//InitDB initialisation db
func InitDB() *sql.DB {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	connStr := "user=" + user + " password=" + password + " host=" + host + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	repo.db = db
	return db
}

//ReadGroups from database
func ReadGroups() models.Groups {
	var resp models.Groups
	var groups []models.Group
	res, err := repo.db.Query("select * from groups")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var group models.Group
		err := res.Scan(&group.GroupID, &group.Title)
		if err != nil {
			panic(err)
		}
		group.Tasks = getGroupByGroupID(group.GroupID)
		groups = append(groups, group)
	}
	resp.Groups = groups
	return resp
}

func getGroupByGroupID(id int) []models.Task {
	var tasks []models.Task
	res, err := repo.db.Query("select * from tasks where groupid = $1", id)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		var task models.Task
		err := res.Scan(&task.TaskID, &task.Title, &task.GroupID)
		if err != nil {
			panic(err)
		}
		task.TimeFrames = getTimeFramesByTaskID(task.TaskID)
		tasks = append(tasks, task)
	}
	return tasks
}

func getTimeFramesByTaskID(id int) []models.TimeFrame {
	var timeframes []models.TimeFrame
	res, err := repo.db.Query("select * from timeframes where taskid = $1", id)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		var timeframe models.TimeFrame
		err := res.Scan(&timeframe.TaskID, &timeframe.From, &timeframe.To)
		if err != nil {
			panic(err)
		}
		timeframes = append(timeframes, timeframe)
	}
	return timeframes
}

//ReadTasks from database
func ReadTasks() models.Tasks {
	var resp models.Tasks
	var tasks []models.Task
	res, err := repo.db.Query("select * from tasks")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var task models.Task
		err := res.Scan(&task.TaskID, &task.Title, &task.GroupID)
		if err != nil {
			panic(err)
		}
		timefr := getTimeFramesByTaskID(task.TaskID)
		task.TimeFrames = timefr
		tasks = append(tasks, task)
	}
	resp.Tasks = tasks
	return resp
}

// CreateTask  add task to database
func CreateTask(task *models.Task) error {
	err := repo.db.QueryRow("INSERT INTO tasks(title, groupid) VALUES ($1,$2) returning taskId", task.Title, task.GroupID).Scan(&task.TaskID)
	return err
}

// CreateGroup - add group to database
func CreateGroup(group *models.Group) error {
	err := repo.db.QueryRow("INSERT INTO groups(title) VALUES ($1) returning groupId", group.Title).Scan(&group.GroupID)
	return err
}

// CreateTimeframe add timeframe to database
func CreateTimeframe(timeframe *models.TimeFrame) error {
	_, err := repo.db.Query("INSERT INTO timeframes(taskId, start, stop) VALUES ($1,$2,$3)", timeframe.TaskID, timeframe.From, timeframe.To)
	return err
}

// UpdateTask update task
func UpdateTask(task *models.Task) error {
	_, err := repo.db.Exec("UPDATE tasks SET title = $1, groupId = $2 where taskId = $3", task.Title, task.GroupID, task.TaskID)
	return err
}

// UpdateGroup update group
func UpdateGroup(group *models.Group) error {
	_, err := repo.db.Exec("UPDATE groups SET title = $1 where groupId = $2", group.Title, group.GroupID)
	return err
}

// DeleteTask delete task by id
func DeleteTask(id int) error {
	_, err := repo.db.Exec("DELETE FROM tasks WHERE taskId = $1", id)
	return err
}

// DeleteGroup delete task by id
func DeleteGroup(id int) error {
	_, err := repo.db.Exec("DELETE FROM groups WHERE groupId = $1", id)
	return err
}

// DeleteTimeframes delete task by id
func DeleteTimeframes(id int) error {
	_, err := repo.db.Exec("DELETE FROM timeframes WHERE taskId = $1", id)
	return err
}
