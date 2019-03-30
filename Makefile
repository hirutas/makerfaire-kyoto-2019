.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=6 go build -o main main.go

.PHONY: deploy
deploy:
	scp main pi@raspberrypi.local:
