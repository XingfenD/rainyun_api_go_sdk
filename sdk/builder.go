package sdk

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

// SDKBuilder constructs a RainyunSDK fluently.
type SDKBuilder struct {
	apiKey string
	trace  *apis.TraceOptions
}

func NewBuilder(apiKey string) *SDKBuilder {
	return &SDKBuilder{apiKey: apiKey}
}

// WithTrace enables structured tracing with the given options. Passing nil
// disables tracing.
func (b *SDKBuilder) WithTrace(options *apis.TraceOptions) *SDKBuilder {
	b.trace = options
	return b
}

func (b *SDKBuilder) Build() *RainyunSDK {
	return newSDK(apis.NewBuilder(b.apiKey).WithTrace(b.trace).Build())
}
