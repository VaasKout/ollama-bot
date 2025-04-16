package ollama_api

import (
	"encoding/json"
	"fmt"
)

type PromptRedisData struct {
	Text      string `json:"text"`
	ChatId    int64  `json:"chat_id"`
	MessageId int64  `json:"message_id"`
}

func (p *PromptRedisData) ToJson() string {
	result, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(result)
}

func PromptDataFromJson(body string) *PromptRedisData {
	var prompt PromptRedisData
	err := json.Unmarshal([]byte(body), &prompt)
	if err != nil {
		fmt.Println(err)
		return &PromptRedisData{}
	}
	return &prompt
}

type GenerateRequest struct {
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	Stream  bool   `json:"stream"`
	Format  string `json:"format"`
	Context []int  `json:"context"`
}

func (r *GenerateRequest) ToJson() []byte {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return result
}

type Answer struct {
	Model     string `json:"model"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
	Timestamp int64  `json:"timestamp"`
	Context   []int  `json:"context"`
}

func AnswerFromJson(data []byte) *Answer {
	var answer Answer
	err := json.Unmarshal(data, &answer)
	if err != nil {
		fmt.Println(err)
		return &Answer{}
	}
	return &answer
}
