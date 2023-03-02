package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

var messages []Message

func main() {
	apiKey := "YOUR_API_KEY"

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ask a question: ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)

		if question == "exit" {
			break
		}

		messages = append(messages, Message{
			Role:    "user",
			Content: question,
		})

		response := getOpenAIResponse(apiKey)
		fmt.Println(response.Choices[0].Messages.Content)
		print("\n")
	}
}

func getOpenAIResponse(apiKey string) OpenaiResponse {
	requestBody := OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	requestJSON, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		println("Error: ", err.Error())
		return OpenaiResponse{}
	}

	messages = append(messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Messages.Content,
	})

	return response
}
