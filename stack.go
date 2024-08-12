package aerrors

import (
	"runtime"
	"strconv"
	"strings"
)

// LogStack return call function stack info from start stack to end stack.
// if end is a positive number, return all call function stack.
func LogStack(start, end int) string {
	var sb strings.Builder
	buf := make([]byte, 0, 32)
	for i := start; i < end || end <= 0; i++ {
		pc, str, line, _ := runtime.Caller(i)
		if line == 0 {
			break
		}
		x := strconv.AppendInt(buf, int64(line), 10)
		sb.Write(StringToBytes(str))
		sb.Write(StringToBytes(":"))
		sb.Write(x)
		sb.Write(StringToBytes("\t"))
		sb.Write(StringToBytes(runtime.FuncForPC(pc).Name()))
		sb.Write(StringToBytes("\n"))
	}
	return sb.String()
}
