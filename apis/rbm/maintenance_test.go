package rbm

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/XingfenD/rainyun_api_go_sdk/apis"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func stubService(t *testing.T, wantMethod, wantPath string, body string) *RbmService {
	t.Helper()
	c := apis.NewRyClient("test-key")
	c.SetTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != wantMethod {
			t.Errorf("method = %s, want %s", req.Method, wantMethod)
		}
		if req.URL.Path != wantPath {
			t.Errorf("path = %s, want %s", req.URL.Path, wantPath)
		}
		return &http.Response{
			StatusCode: 200,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(body)),
		}, nil
	}))
	return NewRbmService(c)
}

func TestGetRbmMaintenance(t *testing.T) {
	svc := stubService(t, "GET", "/product/rbm/789/maintenance", `{"code":200,"data":null}`)
	resp, err := svc.GetRbmMaintenance(789)
	if err != nil {
		t.Fatalf("GetRbmMaintenance() error = %v", err)
	}
	if resp.Code != 200 {
		t.Errorf("code = %d, want 200", resp.Code)
	}
	if resp.Data != nil {
		t.Errorf("data = %v, want nil", resp.Data)
	}
}
