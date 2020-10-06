package logging

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	handlerStart = Event{1, "Handler %s started"}
)

//HandlerStart is an info level log when the handler starts
func HandlerStart(clientName, status, handlerName string, t time.Time) {
	Log.WithFields(logrus.Fields{
		"Handler": handlerName,
		"Status":  status,
		"client":  clientName,
		"Time":    t,
	},
	).Infof(handlerStart.message, handlerName)
}
