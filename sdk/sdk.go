package sdk

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/public"
)

type RainyunSDK struct {
	public.PublicService
}

func New(apiKey string) *RainyunSDK {
	c := apis.NewRyClient(apiKey)
	return &RainyunSDK{
		PublicService: *public.NewPublicService(c),
	}
}
