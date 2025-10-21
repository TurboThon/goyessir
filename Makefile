
build:
	mkdir -p dist
	go build -o dist/goyessir

build-release:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -trimpath -o dist/goyessir_linux_amd64
	GOOS=linux GOARCH=arm64 go build -trimpath -o dist/goyessir_linux_arm64
	GOOS=windows GOARCH=amd64 go build -trimpath -o dist/goyessir_win_amd64.exe

clean:
	rm -rf dist/*
