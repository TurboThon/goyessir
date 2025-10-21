# Goyessir

`goyessir` is a cli tool which aims at easing HTTP interactions.

The tool is especially useful when:
- you need to launch a `python -m http.server` on a host that does not have python
- you need to respond to every requests while logging them
- you need to save files sent by the client

## Usage

```
$ ./goyessir
2025/10/21 22:40:55 http://0.0.0.0:8000/ -> dump request (All methods)
2025/10/21 22:40:55 http://0.0.0.0:8000/f/ -> static filesystem (GET, HEAD)
2025/10/21 22:40:55 http://0.0.0.0:8000/u/ -> file upload (POST, PUT)
2025/10/21 22:40:55 http://0.0.0.0:8000/u/ -> html upload form (GET)
2025/10/21 22:40:55 
  To send files to goyessir, use the following syntax:
  curl http://127.0.0.1:8000/u/ -F "file=@yourfile.txt"
  curl http://127.0.0.1:8000/u/ -F "file[]=@file1.txt" -F "file[]=@file2.txt"

  wget --post-file image.png http://127.0.0.1:8000/u/ -O-

  IWR -Uri http://127.0.0.1:8000/u/ -Method Post -InFile $filePath -UseDefaultCredentials

  Or go to http://127.0.0.1:8000/u/ with a web browser.
```


```
$ ./goyessir -h
Usage of ./goyessir:
  -body-length int
    	Content-Length limit above which the request is saved to a file instead of being printed to stdout (default 8000)
  -c	Enable color output
  -d string
    	Webroot (default ".")
  -debug
    	Enable debug
  -dump-route string
    	Web route where requests are dumped to stdout (default "/")
  -files-route string
    	Web route where the static fs is served (default "/f/")
  -l string
    	Listening address (default "0.0.0.0:8000")
  -log-dir string
    	Directory where requests are saved when they are not printed to stdout (default "requests")
  -no-dirlist
    	Disable directory listing
  -no-upload
    	Disable file upload
  -no-upload-form
    	Disable file upload HTML form
  -upload-dir string
    	Directory where uploaded files are stored (default "uploads")
  -upload-route string
    	Web route to upload files (default "/u/")

```

## Build from source

```
make build-release
```
