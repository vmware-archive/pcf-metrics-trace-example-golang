package logging

import (
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/tracing"
	"os"
	"io"
	"github.com/Sirupsen/logrus"
	"fmt"
)

var logger logrus.StdLogger

func init() {
	logger = logrus.StandardLogger()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)
}

func LogToWriter(out io.Writer) {
	logrus.SetOutput(out)
}

type TraceLogger struct {
	traceLogPrefix string
}

func NewLogger(traceIds *tracing.Ids) *TraceLogger {
	traceLogPrefix := fmt.Sprintf("[%s,%s,%s,true] ",
		traceIds.ParentSpanId,
		traceIds.TraceId,
		traceIds.SpanId,
	)

	return &TraceLogger{
		traceLogPrefix: traceLogPrefix,
	}
}

func (l *TraceLogger) Println(msg string) {
	logger.Println(l.traceLogPrefix + msg)
}

func (l *TraceLogger) Printf(msg string, a ...interface{}) {
	logger.Printf(l.traceLogPrefix + msg, a...)
}
