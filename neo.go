package neo

import (
	"github.com/ivpusic/golog"
	"os"
)

// Type which will be passed as argument of ``panic`` if Neo assertion fails.
type NeoAssertError struct {
	// status which will be returned to user
	status int
	// message which will be returned to user
	message []byte
}

var (
	app *Application
	log = golog.GetLogger("github.com/ivpusic/neo")
)

func init() {
	// default log Level
	log.Level = golog.INFO
}

// Getting Neo Application instance. This is singleton function.
// First time when we call this method function will try to parse configuration for Neo application.
// It will look for configuration file provided by ``--config`` CLI argument (if exist).
func App() *Application {
	if app == nil {
		confFile := ""

		for i, arg := range os.Args {
			if arg == "--config" {
				if len(arg) > i+1 {
					confFile = os.Args[i+1]
					break
				}
			}
		}

		app = &Application{}
		app.init(confFile)
	}

	return app
}

///////////////////////////////////////////////////////////////////
// Neo Assertions
///////////////////////////////////////////////////////////////////

// Asserting some condition. If assertion fails, code bellow assert won't execute, and Neo
// will return to client provided ``status`` and ``message``.
func Assert(condition bool, status int, message []byte) {
	if !condition {
		panic(&NeoAssertError{
			status,
			message,
		})
	}
}

// Same as assert, but instead of passing some boolean condition as first argument, here we are passing object,
// and checking if object is nil.
func AssertNil(obj interface{}, status int, message []byte) {
	if obj != nil {
		Assert(false, status, message)
	}
}

// Same as assert, but instead of passing some boolean condition as first argument, here we are passing object,
// and checking if object is not nil.
func AssertNotNil(obj interface{}, status int, message []byte) {
	if obj == nil {
		Assert(false, status, message)
	}
}
