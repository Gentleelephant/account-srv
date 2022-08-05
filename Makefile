.PHONY: build check help clean update

help:
	@echo "use [make check]: code review"
	@echo "use [make build]: build project"
	@echo "use [make clean]: clean project"
	@echo "use [make update]: git pull project"
	@echo "use [make image]: build docker image. do not use it"
build:
	@go build -o account-srv main.go

check:
	@golangci-lint run

clean:
	@rm -f ./account-srv
	@rm -rf ./cache
	@rm -rf ./logs

update:
	@git pull

image:
	@bash ./build.sh