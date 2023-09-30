.PHONY:

build:
	go build -o ./.bin/bot app/bot.go

run: build
	./.bin/bot