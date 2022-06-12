build:
	cd frontend && npm run build
	go env -w CGO_ENABLED=0
	CGO_ENABLED=0 go build -o bin ./...

run:
	cd frontend && npm run build
	go env -w CGO_ENABLED=0
	CGO_ENABLED=0 go run ./...

release:
	cd frontend && npm run build
	go env -w CGO_ENABLED=0
	CGO_ENABLED=0 go build -ldflags "-s -w" -o bin ./...

releasedocker:
	cd frontend && npm run build
	go env -w CGO_ENABLED=0
	CGO_ENABLED=0 go build -ldflags "-s -w" -o bin ./...
	sudo docker build . -t cybe42/madgo