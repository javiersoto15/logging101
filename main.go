package main

import (
	"time"

	log "github.com/temp/Logging-Tut/logging"
)

func main() {
	log.StartLogger()
	log.HandlerStart("Tutorial", "Started", "Example Handler", time.Now())
}
{"Handler":"Example Handler","Status":"Started","Time":"2020-10-05T20:55:05.982493-05:00","client":"Tutorial","level":"info","msg":"Handler Example Handler started","time":"2020-10-05T20:55:05-05:00"}