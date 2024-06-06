
build:
	mkdir -p dist
	go build -o dist/goyessir

build-release:
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build -o dist/goyessir_linux_amd64
	GOOS=windows GOARCH=amd64 go build -o dist/goyessir_win_amd64.exe

clean:
	rm -rf dist/*
