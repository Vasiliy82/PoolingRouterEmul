package tracer

import (
	"fmt"
	"runtime"
	"strings"
)

func GetSpanNameCallerFn() string {
	fnName := getCallerName(3)
	splitted := strings.Split(fnName, ".")
	if len(splitted) > 2 {
		fnName = fmt.Sprintf("%s @ %s", strings.Join(splitted[(len(splitted)-2):], "."), strings.Join(splitted[:(len(splitted)-2)], "."))
	}
	return fnName
}

func getCallerName(targetFrameIndex int) string {
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}
	return frame.Function
}
