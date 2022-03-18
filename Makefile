.PHONY: service

all: service web

service:
	go build -o bin/promotion-deliver-service cmd/service/main.go

web:
	go build -o bin/promotion-deliver-web cmd/web/main.go
