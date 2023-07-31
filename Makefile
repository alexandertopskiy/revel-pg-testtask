# postgres connection string
# postgres://<username>:<password>@localhost:<port>/<db_name>?sslmode=disable
DB_URL='postgres://postgres:postgres@localhost:5432/test-books?sslmode=disable'

# путь к файлам миграции
MG_PATH='database/migrations'


# Создание новой миграции
# Флаг -seq - для создания последовательных версий
# Пример вызова: make mg_create name=init_mg
mg_create:
	@echo "Создание миграции $(name)!"
	@migrate create -ext sql -dir $(MG_PATH) -seq $(name)
	@echo "Миграция создана!"

# Запуск миграций БД
mg_up:
	@echo "Запуск миграции БД!"
	@migrate -database ${DB_URL} -path $(MG_PATH) up
	@echo "Миграция выполнена!"

# Откат миграций
mg_down:
	@echo "Откат миграции БД!"
	@migrate -database ${DB_URL} -path $(MG_PATH) down
	@echo "Все миграции откачены!"