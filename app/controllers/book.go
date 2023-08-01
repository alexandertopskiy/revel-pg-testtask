package controllers

import (
	"revel-pg-testtask/app/models"
	"strings"
)

// Определение всех запросов для Книг
const (
	getAllBooks = `SELECT book_id, book_title FROM book order by book_id;`

	getAuthorsOfBook = `SELECT author_name FROM book_author JOIN author ON book_author.author_id = author.author_id WHERE book_id = $1;`

	getPublishersOfBook = `SELECT publisher_name FROM book_publisher JOIN publisher ON book_publisher.publisher_id = publisher.publisher_id WHERE book_id = $1;`

	getFilteredBooks = `SELECT DISTINCT B.book_id, B.book_title FROM book B LEFT OUTER JOIN book_author BA ON B.book_id = BA.book_id LEFT OUTER JOIN book_publisher BP ON B.book_id = BP.book_id `
)

// Получение ВСЕХ книг
func allBooks() ([]models.Book, error) {
	books, err := getBooks(getAllBooks)

	return books, err
}

// Получение Книг с условием
func filteredBooks(query string) ([]models.Book, error) {
	books, err := getBooks(getFilteredBooks + query + " order by book_id;")

	return books, err
}

// Общая функция получения Книг с запросом к БД
// (для предотвращения дублирования кода)
func getBooks(query string) ([]models.Book, error) {
	books := []models.Book{}

	rows, err := db.Query(query)

	defer rows.Close()

	if err == nil {
		for rows.Next() {
			var id int
			var title string

			err = rows.Scan(&id, &title)
			if err == nil {
				authors, err := getRelatedDataFor(id, getAuthorsOfBook)
				if err != nil {
					return books, err
				}
				publishers, err := getRelatedDataFor(id, getPublishersOfBook)
				if err != nil {
					return books, err
				}

				currentBook := models.Book{ID: id, Title: title, Author: authors, Publisher: publishers}
				books = append(books, currentBook)
			} else {
				return books, err
			}
		}
	} else {
		return books, err
	}

	return books, err
}

// Получение всех Авторов/Изданий для конкретной книги
// (результат - строка с авторами, разделенными запятой)
func getRelatedDataFor(bookId int, query string) (string, error) {
	// slice (массив) для авторов/издательств книги
	var dataSlice []string
	// строка для объединения всех данных из dataSlice
	var result string

	// делаем запрос к БД на получение данных (авторов/издательств), связанных с данной книгой
	rows, err := db.Query(query, bookId)

	defer rows.Close()

	if err == nil {
		// добавляем строки таблицы в slice
		for rows.Next() {
			var value string
			err = rows.Scan(&value)

			if err == nil {
				dataSlice = append(dataSlice, value)
				result = strings.Join(dataSlice[:], "; ")
			} else {
				return result, err
			}
		}
	} else {
		return result, err
	}

	if result == "" {
		// На странице для пустого поля таблицы показываем прочерк
		result = "-"
	}

	return result, err
}
