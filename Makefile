ifneq (,$(wildcard ./.env))
	include .env
	export
endif

serv:
	go run server/main.go

cli:
	go run client/main.go