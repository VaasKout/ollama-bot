package test

import (
	"fmt"
	"github.com/joho/godotenv"
	"ollama-bot/configs"
	"ollama-bot/internal/ollama_api"
	"testing"
	"time"
)

func TestGenerateResponse(t *testing.T) {

	err := godotenv.Load("../../../configs/.env")
	if err != nil {
		t.Fatal(err)
	}

	config := configs.New()
	ollamaApi := ollama_api.New(config.ModelProps.Model)

	var bodyCh = make(chan ollama_api.Answer)
	var errCh = make(chan error)
	var ticker = time.NewTicker(time.Second)

	text := "Write me a function for reading and writing files using golang"
	context := []int{151644, 6023, 151645, 151648, 271, 151649, 198, 198, 9707, 0, 2585, 646, 358, 7789, 498, 3351, 30, 26525, 232}

	go ollamaApi.GetResponse(text, bodyCh, errCh, context)
	fullMsg := ""

loop:
	for {
		select {
		case answer, ok := <-bodyCh:
			fullMsg += answer.Response
			if !ok {
				fmt.Println("fullMsg:", fullMsg)
				ticker.Stop()
				break loop
			}
		case err = <-errCh:
			if err != nil {
				t.Fatal(err)
			}
		case <-ticker.C:
			fmt.Println(fullMsg)
		}
	}
}
