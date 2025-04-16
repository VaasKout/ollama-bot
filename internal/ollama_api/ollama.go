package ollama_api

import (
	"fmt"
	"net/http"
	"ollama-bot/pkg/core/network"
	"strings"
)

const BaseUrl = "http://127.0.0.1:11434/api/generate"

type OllamaApi interface {
	GetResponse(
		msg string,
		bodyChanel chan Answer,
		errChanel chan error,
		context []int,
	)
}

type OllamaImpl struct {
	model      string
	httpClient network.Client
}

func New(
	model string,
) OllamaApi {
	var httpClient = network.New()
	return &OllamaImpl{
		model:      model,
		httpClient: httpClient,
	}
}

func (p *OllamaImpl) GetResponse(
	msg string,
	bodyChannel chan Answer,
	errChanel chan error,
	context []int,
) {

	var byteChannel = make(chan []byte)

	var request = GenerateRequest{
		Model:   p.model,
		Prompt:  msg,
		Stream:  true,
		Context: context,
	}
	var requestJson = request.ToJson()
	fmt.Println(string(requestJson))
	var httpRequest = &network.HttpRequest{
		Url:    BaseUrl,
		Method: http.MethodPost,
		Body:   requestJson,
	}

	go p.httpClient.MakeStreamRequest(httpRequest, byteChannel, errChanel)

	var writeData = true
loop:
	for {
		select {
		case bytes, ok := <-byteChannel:
			if !ok {
				break loop
			}
			line := string(bytes)
			if strings.Contains(line, "\\u003cthink\\u003e") {
				writeData = false
			}
			if strings.Contains(line, "\\u003c/think\\u003e") {
				writeData = true
				continue
			}
			if writeData {
				fmt.Println(string(bytes))
				var result = AnswerFromJson(bytes)
				if result != nil {
					bodyChannel <- *result
				}
				if result.Done {
					break loop
				}
			}
		}
	}
	close(bodyChannel)
}
