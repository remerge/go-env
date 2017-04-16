package env

import (
	"os"

	"github.com/bobziuchkovski/cue"
)

const Key = "GO_ENV"

var Env string

var log = cue.NewLogger("env")

func init() {
	Set(os.Getenv(Key))
}

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

func IsProd() bool {
	return Env == "production" || Env == "prod"
}

func IsDevel() bool {
	return Env == "development" || Env == "devel" || Env == "dev"
}

func IsTesting() bool {
	return Env == "testing" || Env == "test"
}
