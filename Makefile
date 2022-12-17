BINARY_NAME=jlpt

env:
	cp ./.env.sample ./.env

build:
	go build -o ${BINARY_NAME} src/main.go

run: 
	@docker-compose up -d
	@make build
	./${BINARY_NAME}

clean:
	docker-compose down -v
	rm ${BINARY_NAME}