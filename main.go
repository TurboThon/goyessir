package main

import (
	"io/fs"
	"log"
	"os"

	"github.com/turbothon/goyessir/src/cli"
	"github.com/turbothon/goyessir/src/logger"
	"github.com/turbothon/goyessir/src/router"
	"github.com/turbothon/goyessir/src/types"
)

const HELP_TEXT = `
  To send files to goyessir, use the following syntax:
  curl http://127.0.0.1:8000/u/ -F "file=@yourfile.txt"
  curl http://127.0.0.1:8000/u/ -F "file[]=@file1.txt" -F "file[]=@file2.txt"

  wget --post-file image.png http://127.0.0.1:8000/u/ -O-

  IWR -Uri http://127.0.0.1:8000/u/ -Method Post -InFile $filePath -UseDefaultCredentials

  Or go to http://127.0.0.1:8000/u/ with a web browser.
`

func createDirectories(config *types.Config) {
	permissions := fs.FileMode(0750)
	directories := []string{
		config.WebRoot,
		config.UploadDirectory,
		config.LoggingConfig.RequestLogDirectory,
	}

	for _, directory := range directories {
		err := os.MkdirAll(directory, permissions)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	config := cli.CreateConfig()

	logger.InitLoggers()

	createDirectories(config)

	router := router.CreateRouter(config)

  log.Print(HELP_TEXT)

	router.Run(config.Addr)
}
