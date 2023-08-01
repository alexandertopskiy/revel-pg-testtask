package controllers

import (
	"database/sql"
	"log"

	"github.com/revel/revel"
)

var db *sql.DB

// Инициализация Подключения к Базе данных
func InitDB() {
	// Берем "db.url" из app.conf (или передаем как env-переменную - для докера)
	connstring := revel.Config.StringDefault("db.url", "")
	revel.AppLog.Infof("DB connection string: " + connstring)

	var err error
	db, err = sql.Open("postgres", connstring)

	// Ошибка параметров подключения к БД
	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}
	// Отсутствие подключения к БД (sql.Open НЕ гарантирует установку соединения)
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}

	revel.AppLog.Infof("DB Connected")
}

func init() {
	revel.OnAppStart(InitDB)
}
