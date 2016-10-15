package util

import (
	"fmt"
	"runtime"
	"time"

	"github.com/golang/glog"
)

// logPanic logs the caller tree when a panic occurs.
func logPanic(r interface{}) {
	callers := getCallers(r)
	glog.Errorf("Observed a panic: %#v (%v)\n%v", r, r, callers)
}

func getCallers(r interface{}) string {
	callers := ""
	for i := 0; true; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		callers = callers + fmt.Sprintf("%v:%v\n", file, line)
	}

	return callers
}

func RecoverFromPanic(err *error) {
	if r := recover(); r != nil {
		callers := getCallers(r)

		*err = fmt.Errorf(
			"recovered from panic %q. (err=%v) Call stack:\n%v",
			r,
			*err,
			callers)
	}
}
