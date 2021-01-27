package repository

import (
	"database/sql"

	"github.com/pratamawijaya/Golang101/rest-api/models"
)

func GetTasks(db *sql.DB) models.TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := models.TaskCollection{}

	for rows.Next() {
		task := models.Task{}
		err2 := rows.Scan(&task.ID, &task.Name, &task.Status)

		if err2 != nil {
			panic(err2)
		}

		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func PutTask(db *sql.DB, name string, status int) (int64, error) {
	sql := "INSERT INTO tasks(name, status) VALUES(?,?)"

	// membuat prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit jika error
	if err != nil {
		panic(err)
	}
	// memastikan statement ditutup setelah selesai
	defer stmt.Close()

	result, err2 := stmt.Exec(name, status)
	// Exit jika error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
