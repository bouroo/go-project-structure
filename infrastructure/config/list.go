package config

// KEYS_USED is used to check if all keys are used in config and required
var KEYS_USED = map[string]bool{
	"app.debug":            false,
	"app.listen":           true,
	"app.port.http":        true,
	"app.port.grpc":        true,
	"jwt.key":              true,
	"db.postgres.host":     true,
	"db.postgres.port":     true,
	"db.postgres.user":     true,
	"db.postgres.password": true,
	"db.postgres.database": true,
}
