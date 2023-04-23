build:
	docker-compose build app

run:
	docker-compose up app

migrate:
	migrate -path ./schema -database 'postgres://postgres:postgres:5437/postgres?sslmode=disable' up
