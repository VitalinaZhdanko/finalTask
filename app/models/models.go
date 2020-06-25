package models

import "time"

type All struct {
	Groups []Group `json:"groups"`
	Tasks  []Task  `json:"tasks"`
}

// Groups struct
type Groups struct {
	Groups []Group `json:"groups"`
}

//Tasks struct
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// Group struct
type Group struct {
	GroupID int    `json:"groupId"`
	Title   string `json:"title"`
	Tasks   []Task `json:"tasks,omitempty"`
}

// Task struct
type Task struct {
	TaskID     int         `json:"taskId"`
	Title      string      `json:"title"`
	GroupID    int         `json:"groupId"`
	TimeFrames []TimeFrame `json:"timeFrames,omitempty"`
}

// TimeFrame struct
type TimeFrame struct {
	TaskID int       `json:"taskId"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}

//add validate
