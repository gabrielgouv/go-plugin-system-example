build-plugins:
	go build -buildmode=plugin -o ./plugins/plugin-a/ ./plugins/plugin-a/main.go
	go build -buildmode=plugin -o ./plugins/plugin-b/ ./plugins/plugin-b/main.go

run:
	go run app/main.go

start:
	make build-plugins
	make run