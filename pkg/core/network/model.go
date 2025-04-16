package network

const GET_METHOD = "GET"
const POST_METHOD = "POST"

type HttpRequest struct {
	Url     string
	Body    []byte
	Headers map[string]string
	Method  string
}

type HttpResponse struct {
	Body       []byte
	Error      error
	StatusCode int
}
