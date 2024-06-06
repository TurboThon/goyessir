package router

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/turbothon/goyessir/src/fileupload"
	"github.com/turbothon/goyessir/src/requestlog"
	"github.com/turbothon/goyessir/src/types"
)

func CreateRouter(config *types.Config) *gin.Engine {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.SetTrustedProxies(nil)

	setupDumpRequests(r, config)

	setupStaticFS(r, config)

	if config.FileUpload {
		setupFileUpload(r, config)
	}

	return r
}

func setupStaticFS(engine *gin.Engine, config *types.Config) {
	fs := gin.Dir(config.WebRoot, !config.NoDirListing)
	engine.StaticFS(config.Routes.StaticFS, fs)
	log.Printf("http://%s%s -> static filesystem (GET, HEAD)", config.Addr, config.Routes.StaticFS)
}

func setupFileUpload(engine *gin.Engine, config *types.Config) {
	engine.MaxMultipartMemory = 8 << 22 // 32 MiB, the default

	if !strings.HasSuffix(config.Routes.Upload, "/") {
		config.Routes.Upload = config.Routes.Upload + "/"
	}

	path_wildcard := fmt.Sprintf("%s*filepath", config.Routes.Upload)

	engine.POST(path_wildcard, fileupload.FileUploadHandlerCreate(config))
	engine.PUT(path_wildcard, fileupload.FileUploadHandlerCreate(config))
	log.Printf("http://%s%s -> file upload (POST, PUT)", config.Addr, config.Routes.Upload)
}

func setupDumpRequests(engine *gin.Engine, config *types.Config) {
	// engine.Any(fmt.Sprintf("%s/*route", config.Routes.Dump), requestlog.LogRequestMiddleware(config))
	// engine.Any(config.Routes.Dump, requestlog.LogRequestMiddleware(config))
	engine.Use(requestlog.LogRequestMiddleware(config))

	log.Printf("http://%s%s -> dump request (All methods)", config.Addr, config.Routes.Dump)
}
