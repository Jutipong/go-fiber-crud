## run first project by this command:

```bash
go get
```

## Auto Generate Database(ge-gen)

```bash
go run dal/gen.go
```

## Start Project

- F5 (Debug)
- gomon .

## fieldalignment

```bash
go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
fieldalignment -fix <package_path>
```

## golangci

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
golangci-lint run ./...
```

## Extension vs code

- Go
- Error Lens
- Go Test Explorer
- REST Client
- YAML

## Test Call api

- file \*.http ใช้สำหรับเรียก api
