package constant

type HTTPMethod string

const (
	HTTPMethod_GET    HTTPMethod = "GET"
	HTTPMethod_POST   HTTPMethod = "POST"
	HTTPMethod_DELETE HTTPMethod = "DELETE"
	HTTPMethod_PUT    HTTPMethod = "PUT"
	HTTPMethod_PATCH  HTTPMethod = "PATCH"
)
