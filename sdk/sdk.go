package sdk

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/public"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
)

type RainyunSDK struct {
	public.PublicService
	rcs.RcsService
}

func New(apiKey string) *RainyunSDK {
	c := apis.NewRyClient(apiKey)
	return &RainyunSDK{
		PublicService: *public.NewPublicService(c),
		RcsService:    *rcs.NewRcsService(c),
	}
}
