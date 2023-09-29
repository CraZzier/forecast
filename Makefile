serve:
	go run .
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/weather .
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/weather.exe .