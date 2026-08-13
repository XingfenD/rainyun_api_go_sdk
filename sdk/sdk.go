package sdk

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/domain"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/expense"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/product"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/public"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rbm"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rca"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/ros"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/ssl"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/user"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/workorder"
)

type TraceSink = apis.TraceSink
type TraceOptions = apis.TraceOptions
type HTTPTrace = apis.HTTPTrace
type ResultTrace = apis.ResultTrace

type RainyunSDK struct {
	public.PublicService
	domain.DomainService
	product.ProductService
	rbm.RbmService
	rca.RcaService
	rcs.RcsService
	rgs.RgsService
	ros.RosService
	ssl.SslService
	user.UserService
	workorder.WorkorderService
	expense.ExpenseService
}

func New(apiKey string) *RainyunSDK {
	return newSDK(apis.NewRyClient(apiKey))
}

// NewWithTrace creates an SDK client that emits structured trace events to
// options.Sink.
func NewWithTrace(apiKey string, options TraceOptions) *RainyunSDK {
	return newSDK(apis.NewRyClientWithTrace(apiKey, options))
}

func NewTraceOptions(sink TraceSink) TraceOptions {
	return apis.NewTraceOptions(sink)
}

func newSDK(c *apis.RyClient) *RainyunSDK {
	return &RainyunSDK{
		PublicService:    *public.NewPublicService(c),
		DomainService:    *domain.NewDomainService(c),
		ProductService:   *product.NewProductService(c),
		RbmService:       *rbm.NewRbmService(c),
		RcaService:       *rca.NewRcaService(c),
		RcsService:       *rcs.NewRcsService(c),
		RgsService:       *rgs.NewRgsService(c),
		RosService:       *ros.NewRosService(c),
		SslService:       *ssl.NewSslService(c),
		UserService:      *user.NewUserService(c),
		WorkorderService: *workorder.NewWorkorderService(c),
		ExpenseService:   *expense.NewExpenseService(c),
	}
}
