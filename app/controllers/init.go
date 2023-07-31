package controllers

import (
	"database/sql"
	"fmt"

	"github.com/revel/revel"
)

var db *sql.DB

// Инициализация Подключения к Базе данных
func InitDB() {
	connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", "postgres", "postgres", "test-books")

	var err error
	db, err = sql.Open("postgres", connstring)
	if err != nil {
		revel.AppLog.Infof("DB Error", err)
	}
	revel.AppLog.Infof("DB Connected")
}

func init() {
	revel.OnAppStart(InitDB)
}
