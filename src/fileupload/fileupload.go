package fileupload

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/turbothon/goyessir/src/types"
)

func computeSuffix(file_dest string) string {

	final_file_dest := file_dest
	suffix_number := 0

	_, err := os.Stat(final_file_dest)
	for err == nil || !errors.Is(err, os.ErrNotExist) {
		// File exist, try adding a number until the file does not exist
		suffix_number += 1
		final_file_dest = fmt.Sprintf("%s.%d", file_dest, suffix_number)
		_, err = os.Stat(final_file_dest)
	}
	return final_file_dest
}

func saveFile(c *gin.Context, file *multipart.FileHeader, file_dest string) error {
	final_file_dest := computeSuffix(file_dest)
	err := c.SaveUploadedFile(file, final_file_dest)
	return err
}

func handleMultipartForm(c *gin.Context, form *multipart.Form, config *types.Config) {
	remote_ip := c.RemoteIP()
	if form == nil {
		return
	}
	for _, files := range form.File {
		for _, file := range files {
			if file == nil {
				continue
			}

			// Upload the file to specific dst.
			sanitized_filename := filepath.Clean(filepath.Join("/", file.Filename))
			sanitized_filepath := filepath.Join(config.UploadDirectory, sanitized_filename)
			err := saveFile(c, file, sanitized_filepath)
			if err != nil {
				log.Printf("Error while saving file: %+v", err)
			} else {
				log.Printf("Saved file %s from %s", file.Filename, remote_ip)
			}
		}
	}
}

func handleOtherContentType(c *gin.Context, config *types.Config) {
	body_reader, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error while reading body: %+v", err)
		return
	}

	extensions, err := mime.ExtensionsByType(c.ContentType())
	extension := ".unknown-ext"

	if err == nil && len(extensions) > 0 {
		extension = extensions[0]
	}

	filename := fmt.Sprintf("body-%s%s", time.Now().UTC().Format(time.RFC3339), extension)
	filepath := filepath.Join(config.UploadDirectory, filename)

	err = os.WriteFile(filepath, body_reader, 0640)
	if err != nil {
		log.Printf("Error while writing body to file: %+v", err)
	} else {
		log.Printf("Saved body to file %s", filepath)
	}
}

func FileUploadHandlerCreate(config *types.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err == nil && form != nil {
			handleMultipartForm(c, form, config)
		} else {
			handleOtherContentType(c, config)
		}

		c.String(http.StatusOK, "200 OK")
	}
}
