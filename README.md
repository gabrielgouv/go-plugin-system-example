# go-plugin-system-example

This is an simple example of project using go plugin lib

## Running

### 1. Using Make
You need Make installed in your machine. <br>
To build and run in one shot run: `make start`

To build and run separately: <br>
1. Run `make build`
2. Run `make run`

### 2. Using default
Build using go build:
```
go build -buildmode=plugin -o ./plugins/plugin-a/ ./plugins/plugin-a/main.go
go build -buildmode=plugin -o ./plugins/plugin-b/ ./plugins/plugin-b/main.go
```

Run using go run:
```
go run app/main.go
```

