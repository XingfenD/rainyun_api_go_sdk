package sdk

import (
	"testing"

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

type sinkStub struct{}

func (sinkStub) OnHTTPTrace(HTTPTrace)     {}
func (sinkStub) OnResultTrace(ResultTrace) {}

func assertServicesInitialized(t *testing.T, s *RainyunSDK) {
	t.Helper()
	if s == nil {
		t.Fatal("sdk is nil")
	}
	checks := []struct {
		name string
		ok   bool
	}{
		{"PublicService", s.PublicService != (public.PublicService{})},
		{"DomainService", s.DomainService != (domain.DomainService{})},
		{"ProductService", s.ProductService != (product.ProductService{})},
		{"RbmService", s.RbmService != (rbm.RbmService{})},
		{"RcaService", s.RcaService != (rca.RcaService{})},
		{"RcsService", s.RcsService != (rcs.RcsService{})},
		{"RgsService", s.RgsService != (rgs.RgsService{})},
		{"RosService", s.RosService != (ros.RosService{})},
		{"SslService", s.SslService != (ssl.SslService{})},
		{"UserService", s.UserService != (user.UserService{})},
		{"WorkorderService", s.WorkorderService != (workorder.WorkorderService{})},
		{"ExpenseService", s.ExpenseService != (expense.ExpenseService{})},
	}
	for _, c := range checks {
		if !c.ok {
			t.Errorf("%s not initialized", c.name)
		}
	}
}

func TestNew(t *testing.T) {
	assertServicesInitialized(t, New("test-key"))
}

func TestNewWithTrace(t *testing.T) {
	assertServicesInitialized(t, NewWithTrace("test-key", NewTraceOptions(sinkStub{})))
}
