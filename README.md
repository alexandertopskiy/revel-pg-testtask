# Тестовое задание для Frontend-разработчика на Revel (Go)

__Стек: Revel (Go), PostgreSql__

## Описание задания

### Цель

Продемонстрировать владение определенным стеком технологий:

- Go, Revel
- PostgreSql
- HTML, CSS, JavaScript

### Начальные условия

Представим что у нас есть некий список __книг__, __авторов__ и __издательств__.

- У каждой __книги__ есть __автор__ и __издательство__, в котором она издана.
- У __автора__ может быть __несколько книг__, и он может работать с __разными издательствами__.
- __Книга__ может быть издана сразу __несколькими издательствами__.
- Над некоторыми __книгами__ могли работать __несколько авторов__.
- Наш __список авторов и издательств неполный__ - то есть в базе данных присутствуют записи с __книгами__, которые __не соответствуют какому либо автору и/или издательству__.

### Задача

1. Продумать бизнес-логику базы данных.
2. Создать оптимизированную структуру таблиц под данную задачу.
3. Сверстать страничку, которая будет отображать списки книг, авторов и издательств. Возможно использование css шаблонов (Bootstram, Tailwind и т.д.).
4. Реализовать возможность сортировки данных.
5. Сделать возможность вывода списков в соответствии с заданными исходными данными, например:
      - вывод списка авторов соответствующих определенному издательству
      - вывод списка книг соответствующих определенному автору и/или издательству
      - вывода списка книг, которые не соответствуют ни одному автору, и/или ни одному издательству в нашем списке.

### Инструменты

Для выполнения данного задания необходимо использовать

- В качестве Базы данных: __PostgreSql__
- Фреймворк [__Revel__](http://revel.github.io/tutorial/gettingstarted.html), который написан на языке __Go__.

### Результат

- Результат опубликовать на сайте __github.com__ в виде репозитория
- В процессе работы необходимо активно использовать __git__ и создавать __коммиты__ на основных этапах разработки.
- Бизнес-логику базы данных оформить в виде __миграций__ и опубликовать в __отдельной папке__ проекта.

## Описание решения

## Запуск проекта

### Start the web server

   revel run myapp

### Go to <http://localhost:9000/> and you'll see

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

## Help

- The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
- The [Revel guides](http://revel.github.io/manual/index.html).
- The [Revel sample apps](http://revel.github.io/examples/index.html).
- The [API documentation](https://godoc.org/github.com/revel/revel).
