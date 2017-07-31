package tracing

import "net/http"

var (
	TraceIdHeader = "X-B3-TraceId"
	SpanIdHeader  = "X-B3-SpanId"
	ParentSpanIdHeader = "X-B3-ParentSpanId"
)

type Ids struct {
	TraceId string
	SpanId string
	ParentSpanId string
}

func NewIds(headers http.Header) *Ids {
	return &Ids{
		TraceId: headers.Get(TraceIdHeader),
		SpanId: headers.Get(SpanIdHeader),
		ParentSpanId: headers.Get(ParentSpanIdHeader),
	}
}