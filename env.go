package env

import (
	"os"

	"github.com/remerge/cue"
)

// Key is the OS environment variable that holds the environment name.
const Key = "GO_ENV"

// Env is the Go variable that holds the environment name.
var Env string

var log = cue.NewLogger("env")

func init() {
	Set(os.Getenv(Key))
}

// Set the environment.
func Set(name string) {
	if name == "" {
		name = "development"
	}

	log.WithFields(cue.Fields{
		"old": Env,
		"new": name,
	}).Debug("set")

	if err := os.Setenv(Key, name); err != nil {
		log.Panic(err, "failed to set env")
	} else {
		Env = name
	}
}

// IsProd returns true if the current environment is `production` or `prod`.
func IsProd() bool {
	return Env == "production" || Env == "prod"
}

// IsDevel returns true if the current environment is `development`, `devel` or
// `dev`.
func IsDevel() bool {
	return Env == "development" || Env == "devel" || Env == "dev"
}

// IsTesting returns true if the current environment is `testing` or `test`.
func IsTesting() bool {
	return Env == "testing" || Env == "test"
}
