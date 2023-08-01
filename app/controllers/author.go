package controllers

import (
	"revel-pg-testtask/app/models"
)

// Определение всех запросов для Авторов
const (
	getAllAuthors = `SELECT author_id, author_name FROM author order by author_id;`

	getFilteredAuthors = `SELECT distinct AP.author_id, author_name FROM author_publisher AP JOIN author A ON AP.author_id = A.author_id `
)

// Получение ВСЕХ авторов
func allAuthors() ([]models.Author, error) {
	authors, err := getAuthors(getAllAuthors)

	return authors, err
}

// Получение Авторов с условием
func filteredAuthors(authorQuery string, publisherQuery string) ([]models.Author, error) {
	var query = getFilteredAuthors + makeAuthorCondition(authorQuery, publisherQuery) + " order by AP.author_id;"

	authors, err := getAuthors(query)

	return authors, err
}

// Общая функция получения Авторов с запросом к БД
// (для предотвращения дублирования кода)
func getAuthors(query string) ([]models.Author, error) {
	authors := []models.Author{}

	rows, err := db.Query(query)

	defer rows.Close()

	if err == nil {
		for rows.Next() {
			var id int
			var name string

			err = rows.Scan(&id, &name)
			if err == nil {
				currentAuthor := models.Author{ID: id, Name: name}
				authors = append(authors, currentAuthor)
			} else {
				return authors, err
			}
		}
	} else {
		return authors, err
	}

	return authors, err
}

// Составление полного условия для запроса "Авторы"
func makeAuthorCondition(authorQuery string, publisherQuery string) string {
	var condition string

	if authorQuery != "" && publisherQuery != "" {
		condition = "where A." + authorQuery + " AND " + publisherQuery
	} else if authorQuery != "" {
		condition = "where A." + authorQuery
	} else if publisherQuery != "" {
		condition = "where " + publisherQuery
	}

	return condition
}

// Получение из authorId условия-строки для фильтра по авторам
func getAuthorCondition(authorId string) string {
	var condition string
	if authorId == "null" {
		condition += "author_id is NULL"
	} else if authorId == "" {
		condition += ""
	} else {
		condition += "author_id = " + authorId
	}

	return condition
}
