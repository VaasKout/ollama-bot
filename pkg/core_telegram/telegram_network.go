package core_telegram

import (
	"encoding/json"
	"fmt"
	"ollama-bot/pkg/core/network"
	"ollama-bot/pkg/logger"

	"log"
)

type TelegramNetworkApi interface {
	GetUpdate(offset int64) (*TelegramResponse, error)
	SendMessage(message *OutputMessage) (*TelegramSendResponse, error)
	ProcessCallback(operation CallbackOperations) *network.HttpResponse
}

type TelegramNetworkImpl struct {
	networkApi network.Client
	botKey     string
	logger     *logger.Logger
}

func New(logger *logger.Logger, token string) *TelegramNetworkImpl {
	var networkApi = network.New()

	return &TelegramNetworkImpl{
		networkApi: networkApi,
		botKey:     token,
		logger:     logger,
	}
}

func (telegramNetwork *TelegramNetworkImpl) ProcessCallback(operation CallbackOperations) *network.HttpResponse {
	switch operation.(type) {
	case *EditMessage:
		var result = telegramNetwork.editMessage(operation)
		return result
	case *EditKeyboard:
		var result = telegramNetwork.editKeyboard(operation)
		return result
	}

	return nil
}

func (telegramNetwork *TelegramNetworkImpl) editMessage(editMessageRequest CallbackOperations) *network.HttpResponse {
	var request = telegramNetwork.getHttpRequest(editMessageRequest, EditMessageTextMethod)
	return telegramNetwork.networkApi.MakeRequest(request)
}

func (telegramNetwork *TelegramNetworkImpl) editKeyboard(keyboard CallbackOperations) *network.HttpResponse {
	var request = telegramNetwork.getHttpRequest(
		keyboard,
		EditMessageKeyboard,
	)
	return telegramNetwork.networkApi.MakeRequest(request)
}

func (telegramNetwork *TelegramNetworkImpl) SendMessage(message *OutputMessage) (*TelegramSendResponse, error) {
	switch {
	case message.PhotoId != "":
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendPhotoMethod))
	case message.VideoId != "":
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendVideoMethod))
	case message.VoiceId != "":
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendVoiceMethod))
	case message.StickerId != "":
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendStickerMethod))
	case message.AudioId != "":
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendAudioMethod))
	case message.DocumentId != "":
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendDocumentMethod))
	default:
		return telegramNetwork.processRequest(telegramNetwork.getHttpRequest(message, SendMessageMethod))
	}
}

func (telegramNetwork *TelegramNetworkImpl) GetUpdate(offset int64) (*TelegramResponse, error) {
	var url = fmt.Sprintf(
		"https://api.telegram.org/bot%s/getUpdates?offset=%d",
		telegramNetwork.botKey,
		offset,
	)
	var request = &network.HttpRequest{
		Url:    url,
		Method: network.GET_METHOD,
	}

	result := telegramNetwork.networkApi.MakeRequest(request)
	if result.StatusCode != 200 {
		telegramNetwork.logger.Error("BODY: ", string(result.Body), " ERROR: ", result.Error.Error())
		return &TelegramResponse{}, result.Error
	}
	var telegramResponse TelegramResponse
	err := json.Unmarshal(result.Body, &telegramResponse)
	if err != nil {
		log.Println(err.Error())
	}
	return &telegramResponse, err

}

func (telegramNetwork *TelegramNetworkImpl) processRequest(request *network.HttpRequest) (*TelegramSendResponse, error) {
	result := telegramNetwork.networkApi.MakeRequest(request)
	if result.StatusCode != 200 {
		telegramNetwork.logger.Error("BODY: ", string(result.Body), " ERROR: ", result.Error.Error())
		return &TelegramSendResponse{}, result.Error
	}
	fmt.Println(string(result.Body))
	var telegramResponse TelegramSendResponse
	err := json.Unmarshal(result.Body, &telegramResponse)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &telegramResponse, err
}

func (telegramNetwork *TelegramNetworkImpl) validateResponse(response *network.HttpResponse) {
	var errorDescription ErrorMessage

	if response.StatusCode == 200 {
		return
	}

	telegramNetwork.logger.Error(
		"BODY: ",
		string(response.Body),
		" CODE: ",
		response.StatusCode,
		" ERROR: ",
		response.Error,
	)

	err := json.Unmarshal(response.Body, &errorDescription)
	if err != nil {
		telegramNetwork.logger.Error(err.Error())
		return
	}

	switch errorDescription.Description {
	case VoiceMessagesForbiddenError:
		response.Body = []byte(VoiceMessagesForbiddenMessage)
	}
}

func (telegramNetwork *TelegramNetworkImpl) getHttpRequest(body interface{}, method string) *network.HttpRequest {

	var url = fmt.Sprintf(
		"https://api.telegram.org/bot%s/%s",
		telegramNetwork.botKey,
		method,
	)

	jsonMessage, err := json.Marshal(&body)
	if err != nil {
		telegramNetwork.logger.Error(fmt.Sprintf("Cannot marshal body: %v", body))
	}

	return &network.HttpRequest{
		Url:    url,
		Body:   jsonMessage,
		Method: network.GET_METHOD,
	}
}
