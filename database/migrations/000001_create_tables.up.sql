-- Создание Таблицы 'Книги'
CREATE TABLE
	IF NOT EXISTS book (
		book_id serial NOT NULL,
		book_title VARCHAR(400),
		CONSTRAINT pk_book PRIMARY KEY (book_id)
	);

-- Создание Таблицы 'Авторы'
CREATE TABLE
	IF NOT EXISTS author (
		author_id serial NOT NULL,
		author_name VARCHAR(400),
		CONSTRAINT pk_author PRIMARY KEY (author_id)
	);

-- Создание Таблицы 'Издательства'
CREATE TABLE
	IF NOT EXISTS publisher (
		publisher_id serial NOT NULL,
		publisher_name VARCHAR(400),
		CONSTRAINT pk_publisher PRIMARY KEY (publisher_id)
	);