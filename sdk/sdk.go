package sdk

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/domain"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/public"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rbm"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rca"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcdn"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/ros"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rvh"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/ssl"
)

type RainyunSDK struct {
	public.PublicService
	domain.DomainService
	rbm.RbmService
	rcs.RcsService
	rgs.RgsService
	ros.RosService
	ssl.SslService
	rcdn.RcdnService
	rca.RcaService
	rvh.RvhService

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
		PublicService: *public.NewPublicService(c),
		DomainService: *domain.NewDomainService(c),
		RbmService:    *rbm.NewRbmService(c),
		RcsService:    *rcs.NewRcsService(c),
		RgsService:    *rgs.NewRgsService(c),
		RosService:    *ros.NewRosService(c),
		SslService:    *ssl.NewSslService(c),
		RcdnService:   *rcdn.NewRcdnService(c),
		RcaService:    *rca.NewRcaService(c),
		RvhService:    *rvh.NewRvhService(c),
		client:        c,
	}
}
