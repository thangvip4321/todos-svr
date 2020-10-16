package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Task struct {
	ID   int64
	Task string
	Note string
	Done bool
	MetaData
}

type MetaData struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	AssigneeID int64
}

type TaskHandler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *TaskHandler {
	handler := TaskHandler{db}
	return &handler
}
func (handler *TaskHandler) CreateTask(task Task) (error, *Task) {
	stmt, err := handler.db.Prepare("INSERT INTO task (Task,Note,Done,AssigneeID) VALUES (?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(task.Task, task.Note, task.Done, task.AssigneeID)
	if err != nil {
		return err, nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err, nil
	}
	rows, err := handler.db.Query("SELECT created_at,updated_at FROM task WHERE ID = ?", id)
	if err != nil {
		return err, nil
	}
	rows.Scan(&task.CreatedAt, &task.UpdatedAt)
	task.ID = id
	return nil, &task
}

func (handler *TaskHandler) GetTask(id int) (error, *Task) {
	rows, err := handler.db.Query("SELECT * FROM task WHERE ID = ?", id) // ????????????????
	task := Task{}
	defer rows.Close()
	rows.Next()

	// db.Query didnt work but db.QueryRow worked like a charm here???????
	// need explanation! (db.Query need Next() call to Rows)
	if err != nil {
		return err, nil
	}
	//fmt.Println(rows.Columns())
	fmt.Println(rows)
	rows.Scan(&task.ID, &task.AssigneeID, &task.CreatedAt, &task.UpdatedAt, &task.Task, &task.Note, &task.Done)
	return nil, &task

}

func (handler *TaskHandler) UpdateTask(id int, task *Task) (error, *Task) {
	rows, err := handler.db.Query("SELECT * FROM task where ID = ?", id)
	if exist := rows.Next(); exist == false {
		return errors.New("no such task exist"), nil
	}
	//columns, err := rows.Columns()
	stmt, err := handler.db.Prepare(fmt.Sprintf("UPDATE tasks SET AssigneeID=?, Task=?, Note=?, Done=?, where ID = ?"))
	defer rows.Close()
	defer stmt.Close()
	if err != nil {
		return err, nil
	}
	_, err = stmt.Exec(task.AssigneeID, task.Task, task.Note, task.Done, id)
	if err != nil {
		return err, nil
	}
	newTask := Task{}
	rows.Next()
	rows.Scan(&newTask.ID, &newTask.AssigneeID, &newTask.CreatedAt, &newTask.UpdatedAt, newTask.Task, &newTask.Note, &newTask.Done)
	return nil, &newTask
}

func (handler *TaskHandler) DeleteTask(id int, task *Task) (error, string) {
	rows, err := handler.db.Query("SELECT * FROM task where ID = ?", id)
	if exist := rows.Next(); exist == false {
		return errors.New("no such task exist"), ""
	}
	rows1, err := handler.db.Query("DELETE FROM tasks where ID = ?", id)
	defer rows.Close()
	defer rows1.Close()
	if err != nil {
		return err, ""
	}
	return nil, "deleted"
}
