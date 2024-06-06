
build:
	mkdir -p dist
	go build -o dist/goyessir

clean:
	rm -rf dist/*
