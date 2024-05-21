package logrus_starter

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	red     = "\033[31m"
	cyan    = "\033[36m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	white   = "\033[37m"
	green   = "\033[32m"
	magenta = "\033[35m"
	reset   = "\033[0m"
)

type LogFormat struct{}

func (l *LogFormat) Format(entry *log.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	sprintf := fmt.Sprintf("====> [%s] %s[%s]%s %s \n", timestamp, getLevelColor(entry.Level), strings.ToUpper(entry.Level.String()), reset, entry.Message)
	return []byte(sprintf), nil
}

func getMethodColor(method string) string {
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

func getStatusColor(code int) string {
	switch {
	case code == 0 || code == 200:
		return green
	default:
		return red
	}
}

func getLevelColor(level log.Level) string {
	switch level {
	case log.DebugLevel, log.TraceLevel:
		return cyan
	case log.ErrorLevel, log.FatalLevel:
		return red
	case log.WarnLevel:
		return yellow
	default:
		return green
	}
}
