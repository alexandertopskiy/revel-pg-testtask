package controllers

import (
	"revel-pg-testtask/app/models"
)

// Определение всех запросов для Книг
const (
	getAllPublishers = `SELECT publisher_id, publisher_name FROM publisher order by publisher_id;`

	getFilteredPublishers = `SELECT distinct AP.publisher_id, publisher_name FROM author_publisher AP JOIN publisher P ON AP.publisher_id = P.publisher_id `
)

// Получение ВСЕХ Издательств
func allPublishers() ([]models.Publisher, error) {
	publishers, err := getPublishers(getAllPublishers)

	return publishers, err
}

// Получение Издательств с условием
func filteredPublishers(authorQuery string, publisherQuery string) ([]models.Publisher, error) {
	var query = getFilteredPublishers + makePublisherCondition(authorQuery, publisherQuery) + " order by publisher_id;"

	publishers, err := getPublishers(query)

	return publishers, err
}

// Общая функция получения Издательств с запросом к БД
// (для предотвращения дублирования кода)
func getPublishers(query string) ([]models.Publisher, error) {
	publishers := []models.Publisher{}

	rows, err := db.Query(query)

	defer rows.Close()

	if err == nil {
		for rows.Next() {
			var id int
			var name string

			err = rows.Scan(&id, &name)
			if err == nil {
				currentPublisher := models.Publisher{ID: id, Name: name}
				publishers = append(publishers, currentPublisher)
			} else {
				return publishers, err
			}
		}
	} else {
		return publishers, err
	}

	return publishers, err
}

// Составление полного условия для запроса "Издательства"
func makePublisherCondition(authorQuery string, publisherQuery string) string {
	var condition string
	if authorQuery != "" && publisherQuery != "" {
		condition = "where " + authorQuery + " AND P." + publisherQuery
	} else if authorQuery != "" {
		condition = "where " + authorQuery
	} else if publisherQuery != "" {
		condition = "where P." + publisherQuery
	}

	return condition
}

// Получение из publisherId условия-строки для фильтра по издательствам
func getPublisherCondition(publisherId string) string {
	var condition string
	if publisherId == "null" {
		condition += "publisher_id is NULL"
	} else if publisherId == "" {
		condition += ""
	} else {
		condition += "publisher_id = " + publisherId
	}

	return condition
}
