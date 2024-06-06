package cli

import (
	"flag"

	"github.com/turbothon/goyessir/src/types"
)

func CreateConfig() *types.Config {
	config := &types.Config{}

	flag.BoolVar(&config.Debug, "debug", false, "Enable debug")
	flag.StringVar(&config.Addr, "l", "0.0.0.0:8000", "Listening address")
	flag.StringVar(&config.WebRoot, "d", ".", "Webroot")
	flag.StringVar(&config.UploadDirectory, "upload-dir", "uploads", "Directory where uploaded files are stored")
	flag.BoolVar(&config.NoDirListing, "no-dirlist", false, "Disable directory listing")
	flag.BoolVar(&config.FileUpload, "u", false, "Enable file upload")
	flag.BoolVar(&config.Color, "c", false, "Enable color output")

	flag.StringVar(&config.LoggingConfig.RequestLogDirectory, "log-dir", "requests", "Directory where requests are saved when they are not printed to stdout")
	flag.Int64Var(&config.LoggingConfig.LogBodyLengthLimit, "body-length", 8000, "Content-Length limit above which the request is saved to a file instead of being printed to stdout")

	flag.StringVar(&config.Routes.StaticFS, "files-route", "/", "Web route where the static fs is served")
	flag.StringVar(&config.Routes.Upload, "upload-route", "/", "Web route to upload files")
	flag.StringVar(&config.Routes.Dump, "dump-route", "/", "Web route where requests are dumped to stdout")

	flag.Parse()

	return config
}
