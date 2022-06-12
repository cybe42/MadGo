build:
	cd frontend && npm run build
	go build -o bin ./...

run:
	cd frontend && npm run build
	go run ./...

release:
	cd frontend && npm run build
	go build -ldflags "-s -w" -o bin ./...