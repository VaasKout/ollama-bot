package network

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client interface {
	MakeRequest(r *HttpRequest) *HttpResponse
	MakeStreamRequest(r *HttpRequest, bodyChanel chan<- []byte, errChanel chan<- error)
}

type ClientImpl struct {
	api *http.Client
}

func New() Client {
	api := &http.Client{
		Timeout: time.Second * 30,
	}

	return &ClientImpl{
		api: api,
	}
}

func (client *ClientImpl) MakeStreamRequest(req *HttpRequest, byteChannel chan<- []byte, errChanel chan<- error) {
	if req == nil {
		errChanel <- errors.New("empty request")
		return
	}
	request, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer(req.Body))
	if err != nil {
		errChanel <- errors.New("error due creating request: " + err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")
	if req.Headers != nil {
		for k, v := range req.Headers {
			request.Header.Set(k, v)
		}
	}
	response, err := client.api.Do(request)
	if err != nil {
		errChanel <- errors.New("error due executing request: " + err.Error())
		return
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		byteChannel <- scanner.Bytes()
	}
	close(byteChannel)
	close(errChanel)
}

func (client *ClientImpl) MakeRequest(req *HttpRequest) *HttpResponse {

	if req == nil {
		return &HttpResponse{Error: errors.New("empty request")}
	}

	request, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer(req.Body))
	if err != nil {
		return &HttpResponse{Error: errors.New("error due creating request: " + err.Error())}
	}

	request.Header.Set("Content-Type", "application/json")
	if req.Headers != nil {
		for k, v := range req.Headers {
			request.Header.Set(k, v)
		}
	}

	response, err := client.api.Do(request)
	if err != nil {
		return &HttpResponse{Error: errors.New("error due executing request: " + err.Error())}
	}

	return client.handleBody(response)
}

func (client *ClientImpl) handleBody(response *http.Response) *HttpResponse {
	var httpResponse = new(HttpResponse)
	if response == nil {
		httpResponse.Error = errors.New("empty response")
		return httpResponse
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		httpResponse.Error = errors.New("error due parsing body: " + err.Error())
		return httpResponse
	}

	httpResponse.StatusCode = response.StatusCode
	httpResponse.Body = body
	if httpResponse.StatusCode != 200 {
		httpResponse.Error = errors.New(
			fmt.Sprintf("Not success. StatusCode code: %d. Body: %v", response.StatusCode, string(body)),
		)
	}

	return httpResponse
}
