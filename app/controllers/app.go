package controllers

import (
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// Полные списки книг/авторов/издательств
func (c App) Index() revel.Result {
	// Получаем все книги
	books, err := allBooks()
	if err != nil {
		panic(err)
	}

	// Получаем всех авторов
	authors, err := allAuthors()
	if err != nil {
		panic(err)
	}

	// Получаем все издательства
	publishers, err := allPublishers()
	if err != nil {
		panic(err)
	}

	// Отрисовываем страничку
	return c.Render(books, authors, publishers)
}

// Отфильтрованные списки книг/авторов/издательств
func (c App) FilterLists() revel.Result {
	// получаем из формы id автора (если указан)
	authorId := c.Params.Get("author")
	// составляем условие для фильтра по авторам
	var authorCondition string
	if authorId == "null" {
		authorCondition += "author_id is NULL"
	} else if authorId == "" {
		authorCondition += ""
	} else {
		authorCondition += "author_id = " + authorId
	}

	// получаем из формы id издательства (если указано)
	publisherId := c.Params.Get("publisher")
	// составляем условие для фильтра по издательствам
	var publisherCondition string
	if publisherId == "null" {
		publisherCondition += "publisher_id is NULL"
	} else if publisherId == "" {
		publisherCondition += ""
	} else {
		publisherCondition += "publisher_id = " + publisherId
	}

	// составляем полное условие фильтрации
	var fullCondition string
	if authorCondition != "" && publisherCondition != "" {
		fullCondition = "where " + authorCondition + " AND " + publisherCondition
	} else if authorCondition != "" {
		fullCondition = "where " + authorCondition
	} else if publisherCondition != "" {
		fullCondition = "where " + publisherCondition
	}

	// Получаем отфильтрованные книги
	books, err := filteredBooks(fullCondition)
	if err != nil {
		panic(err)
	}

	// Получаем отфильтрованных авторов
	authors, err := filteredAuthors(authorCondition, publisherCondition)
	if err != nil {
		panic(err)
	}

	// Получаем отфильтрованные издательства
	publishers, err := filteredPublishers(authorCondition, publisherCondition)
	if err != nil {
		panic(err)
	}

	// доп.аргумент для отрисовки кнопки "показать все" в подмодуле filter.html
	filter := true
	// Отрисовываем страничку
	return c.Render(books, authors, publishers, filter, authorId, publisherId)
}
