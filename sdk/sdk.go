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

	client *apis.RyClient
}

func New(apiKey string) *RainyunSDK {
	return newSDK(apis.NewRyClient(apiKey))
}

// RawResponseBody returns the raw response body of the most recent API request.
func (s *RainyunSDK) RawResponseBody() []byte {
	return s.client.RawBody()
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
		client:           c,
	}
}
