package main

import (
	"time"

	"github.com/pkg/errors"
	log "github.com/temp/Logging-Tut/logging"
)

func main() {
	t := time.Now()
	log.StartLogger()
	log.HandlerStart("Tutorial", "Started", "Example Handler", t)
	log.HandlerCompleted("Tutorial", "Finished", "Example Handler", 200, t)
	err := errors.Errorf("Example error")
	log.HandlerError("Tutorial", "Failed", "Example Handler", err, t)
}
