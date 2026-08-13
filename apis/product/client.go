package product

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type ProductService struct {
	client *apis.RyClient
}

func NewProductService(c *apis.RyClient) *ProductService {
	return &ProductService{client: c}
}
