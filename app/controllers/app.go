package controllers

import (
	"revel-pg-testtask/app/models"

	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// Полные списки книг/авторов/издательств
func (c App) Index() revel.Result {
	// Получаем все книги
	books, authors, publishers := getAllData()

	// Отрисовываем страничку
	return c.Render(books, authors, publishers)
}

// Отфильтрованные списки книг/авторов/издательств
func (c App) FilterLists() revel.Result {
	// получаем из формы id автора (если указан)
	authorId := c.Params.Get("author")
	// получаем sql-условие для фильтра по авторам
	authorCondition := getAuthorCondition(authorId)

	// получаем из формы id издательства (если указано)
	publisherId := c.Params.Get("publisher")
	// получаем sql-условие для фильтра по издательствам
	publisherCondition := getPublisherCondition(publisherId)

	// составляем полное условие фильтрации
	fullCondition := makeFullCondition(authorCondition, publisherCondition)

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

	// Получаем полные списки авторов/издательств (для корректного отображения выделенного значения в селектах)
	_, allAuthors, allPublishers := getAllData()

	// доп.аргумент для отрисовки кнопки "показать все" в подмодуле filter.html
	filter := true
	// Отрисовываем страничку
	return c.Render(books, authors, publishers, filter, authorId, publisherId, allAuthors, allPublishers)
}

// Получение полных списков
// Для FilterLists отсюда можно взять полные списки Авторов/Издателей для корректного вывода выбранного фильтра в селектах
func getAllData() ([]models.Book, []models.Author, []models.Publisher) {
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

	return books, authors, publishers
}
