build:
	go build -o bin/main main.go

run:
	go run main.go

dev:
	./bin/air

debug:
	./bin/air -d

version:
	./bin/air -v