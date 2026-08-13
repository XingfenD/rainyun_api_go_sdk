package sdk

// SDKBuilder constructs a RainyunSDK fluently.
type SDKBuilder struct {
	apiKey string
	trace  *TraceOptions
}

func NewBuilder(apiKey string) *SDKBuilder {
	return &SDKBuilder{apiKey: apiKey}
}

// WithTrace enables structured tracing with the given options. Passing nil
// disables tracing.
func (b *SDKBuilder) WithTrace(options *TraceOptions) *SDKBuilder {
	b.trace = options
	return b
}

func (b *SDKBuilder) Build() *RainyunSDK {
	if b.trace == nil {
		return New(b.apiKey)
	}
	return NewWithTrace(b.apiKey, *b.trace)
}
