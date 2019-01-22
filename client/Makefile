run:
	@ go run main.go

test:
	@ go test -cover  -failfast ./...

lint:
	@ golangci-lint run --config=.golangci-lint.yml  ./...

up:
	@ docker-compose up -d

down:
	@ docker-compose down
	@ docker-compose rm