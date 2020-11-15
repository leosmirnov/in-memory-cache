package constants

import "time"

const (
	// Default values for api section in config.
	DefaultAPIHost = "0.0.0.0"
	DefaultAPIPort = 9020

	// http
	ContentType         = "Content-Type"
	ApplicationJSONUTF8 = "application/json; charset=utf-8"
	ApplicationJSON     = "application/json"
	TextJSON            = "text/json"

	// Conf defaults
	DefaultCleanupInterval = 500 * time.Millisecond
	DefaultStopTimeout     = 30 * time.Second
)
