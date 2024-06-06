package requestlog

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/turbothon/goyessir/src/logger"
	"github.com/turbothon/goyessir/src/types"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func SaveRequest(request *http.Request, config *types.Config) (string, error) {
	filename := fmt.Sprintf("%s.req", time.Now().UTC().Format(time.RFC3339Nano))
	filepath := filepath.Join(config.LoggingConfig.RequestLogDirectory, filename)
	request_bytes, err := httputil.DumpRequest(request, true)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(filepath, request_bytes, 0600)
	return filepath, err
}

func PrintRequest(request *http.Request) {
	request_bytes, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Printf("Error while printing request: %+v", err)
	}

	fmt.Println(BytesToTerminalString(request_bytes))
}

func PrintHeader(request *http.Request) {
	request_bytes, err := httputil.DumpRequest(request, false)
	if err != nil {
		log.Printf("Error while processing request: %+v", err)
	}

	fmt.Println(BytesToTerminalString(request_bytes))
}

func BytesToTerminalString(bytes []byte) string {
	text := strings.Map(func(r rune) rune {
		if r == '\n' {
			return r
		}
		if unicode.In(r, unicode.C) {
			return ' '
		}
		return r
	}, string(bytes))

	return text
}

func methodColor(method string, config *types.Config) string {

	if config.Color {
		return reset
	}
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

func LogRequestMiddleware(config *types.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Printf("%13v | %s %-7s %s | %#v",
			c.ClientIP(),
			methodColor(c.Request.Method, config), c.Request.Method, reset,
			c.Request.URL.Path,
		)

		if c.Request.ContentLength < config.LoggingConfig.LogBodyLengthLimit {
			PrintRequest(c.Request)
		} else {
			PrintHeader(c.Request)
		}

		c.Next()

		_, err := SaveRequest(c.Request, config)
		if err != nil {
			logger.ErrorLogger.Print(err)
		}
	}
}
