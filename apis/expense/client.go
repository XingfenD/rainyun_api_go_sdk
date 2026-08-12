package expense

import "github.com/XingfenD/rainyun_api_go_sdk/apis"

type ExpenseService struct {
	client *apis.RyClient
}

func NewExpenseService(c *apis.RyClient) *ExpenseService {
	return &ExpenseService{client: c}
}
