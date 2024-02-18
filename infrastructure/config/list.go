package config

// KEYS_USED is used to check if all keys are used in config and required
var KEYS_USED = map[string]bool{
	"app.debug":     false,
	"app.listen":    true,
	"app.port.http": true,
}
