-- Создание Таблицы 'Книги-Авторы'
CREATE TABLE
	IF NOT EXISTS book_author (
		book_id INT,
		author_id INT,
		CONSTRAINT pk_bookauthor PRIMARY KEY (book_id, author_id),
		CONSTRAINT fk_ba_book FOREIGN KEY (book_id) REFERENCES book (book_id),
		CONSTRAINT fk_ba_author FOREIGN KEY (author_id) REFERENCES author (author_id)
	);

-- Создание Таблицы 'Книги-Издательства'
CREATE TABLE
	IF NOT EXISTS book_publisher (
		book_id INT,
		publisher_id INT,
		CONSTRAINT pk_bookpublisher PRIMARY KEY (book_id, publisher_id),
		CONSTRAINT fk_bp_book FOREIGN KEY (book_id) REFERENCES book (book_id),
		CONSTRAINT fk_bp_publisher FOREIGN KEY (publisher_id) REFERENCES publisher (publisher_id)
	);

-- Создание Таблицы 'Авторы-Издательства'
CREATE TABLE
	IF NOT EXISTS author_publisher (
		author_id INT,
		publisher_id INT,
		CONSTRAINT pk_authorpublisher PRIMARY KEY (author_id, publisher_id),
		CONSTRAINT fk_ap_author FOREIGN KEY (author_id) REFERENCES author (author_id),
		CONSTRAINT fk_ap_publisher FOREIGN KEY (publisher_id) REFERENCES publisher (publisher_id)
	);