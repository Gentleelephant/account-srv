.PHONY: build check help

help:
	@echo "use [make check]: code review"
	@echo "use [make build]: build project"
build:
	@go build -o account-srv main.go

check:
	@golangci-lint run

clean:
	@rm -f ./account-srv
	@rm -rf ./cache
	@rm -rf ./logs