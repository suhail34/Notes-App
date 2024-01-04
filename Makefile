server-local:
	go run cmd/notes-api/main.go
.PHONY: server-local

server-container-start:
	docker run -p 8080:8080 -d --name notes-api suhail12/notes-api:latest
.PHONY: server-container-start

server-container-stop:
	docker container stop notes-api
	docker rm -f notes-api
.PHONY: server-container-stop
