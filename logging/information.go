package logging

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	handlerStart     = Event{1, "Handler %s started"}
	handlerCompleted = Event{1, "Handler %s completed"}
	handlerError     = Event{3, "Handler %s has failed"}
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

//HandlerCompleted is an info level log when the handler is completed
func HandlerCompleted(clientName, status, handlerName string, code int, t time.Time) {
	duration := time.Since(t)
	Log.WithFields(logrus.Fields{
		"Handler":       handlerName,
		"Status":        status,
		"client":        clientName,
		"Time":          duration,
		"Response Code": code,
	},
	).Infof(handlerCompleted.message, handlerName)
}

//HandlerError is an error level log when the handler fails
func HandlerError(clientName, status, handlerName string, err error, t time.Time) {
	Log.WithFields(logrus.Fields{
		"Handler": handlerName,
		"Status":  status,
		"client":  clientName,
		"Time":    t,
		"Error":   err.Error(),
	},
	).Errorf(handlerError.message, handlerName)
}
