package main

import (
	"github.com/bobziuchkovski/cue"
	"github.com/bobziuchkovski/cue/collector"
	env "github.com/remerge/go-env"
)

var log = cue.NewLogger("main")

func main() {
	cue.Collect(cue.DEBUG, collector.Terminal{}.New())
	log.WithValue("env", env.Env).Info("current")
	env.Set("something")
	log.WithValue("env", env.Env).Info("current")
}
