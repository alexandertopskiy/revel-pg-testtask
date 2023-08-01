package controllers

import (
	"database/sql"

	"github.com/revel/revel"
)

var db *sql.DB

// Инициализация Подключения к Базе данных
func InitDB() {
	// Берем "db.url" из app.conf (или передаем как env-переменную для докера)
	connstring := revel.Config.StringDefault("db.url", "")
	revel.AppLog.Infof("DB connstring: ", connstring)

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
