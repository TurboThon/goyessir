package types

type Config = struct {
	Debug           bool
	Addr            string
	WebRoot         string
	UploadDirectory string
	NoDirListing    bool
	FileUpload      bool
	Routes          RoutesConfig
	LoggingConfig   LoggingConfig
	Color           bool
}

type LoggingConfig struct {
	LogBodyLengthLimit  int64
	RequestLogDirectory string
}

type RoutesConfig struct {
	StaticFS string
	Upload   string
	Dump     string
}
