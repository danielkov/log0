package log0

import (
	"runtime"
)

func diff(a, b []level) []level {
	mb := map[level]bool{}
	for _, x := range b {
		mb[x] = true
	}
	ab := []level{}
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

func trace(skip int) (f runtime.Frame) {
	pc := make([]uintptr, 10)
	runtime.Callers(skip, pc)
	f, _ = runtime.CallersFrames(pc[:]).Next()
	return f
}

func contains(s []level, e level) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
