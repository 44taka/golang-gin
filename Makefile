destory:
	docker-compose down --rmi all --volumes --remove-orphans

build:
	docker-compose build

up:
	docker-compose up -d

test:
	docker-compose exec app go test -v ./...