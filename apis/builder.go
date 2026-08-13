package apis

// RyClientBuilder constructs a RyClient fluently.
type RyClientBuilder struct {
	apiKey string
	trace  *TraceOptions
}

func NewBuilder(apiKey string) *RyClientBuilder {
	return &RyClientBuilder{apiKey: apiKey}
}

// WithTrace enables structured tracing with the given options. Passing nil
// disables tracing.
func (b *RyClientBuilder) WithTrace(options *TraceOptions) *RyClientBuilder {
	b.trace = options
	return b
}

func (b *RyClientBuilder) Build() *RyClient {
	if b.trace == nil {
		return NewRyClient(b.apiKey)
	}
	c := newRyClient(b.apiKey)
	c.traceSink = b.trace.Sink
	c.tracePolicy = b.trace.previewPolicy()
	return c
}
