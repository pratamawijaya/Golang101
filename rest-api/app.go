package main

import (
	"database/sql"

	"github.com/pratamawijaya/Golang101/rest-api/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	// init db
	db := initDB("tasks.db")
	migrate(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/task", handlers.PutTask(db))

	e.Logger.Fatal(e.Start(":8989"))

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    		CREATE TABLE IF NOT EXISTS tasks(
        	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		status INTEGER );
    		`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
