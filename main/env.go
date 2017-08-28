package main

import (
	"github.com/remerge/cue"
	"github.com/remerge/cue/collector"
	env "github.com/remerge/go-env"
)

var log = cue.NewLogger("main")

func main() {
	cue.Collect(cue.DEBUG, collector.Terminal{}.New())
	log.WithValue("env", env.Env).Info("current")
	env.Set("something")
	log.WithValue("env", env.Env).Info("current")
}
