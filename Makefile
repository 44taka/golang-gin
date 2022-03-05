# TODO:makefileを整備する
destroy:
	docker-compose down --rmi all --volumes --remove-orphans

build:
	docker-compose build

up:
	docker-compose up -d

test:
	docker-compose exec app bash -c "go clean -testcache && go test -v ./..."