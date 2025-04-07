//go:build integration

package main

import (
	"context"
	"testing"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

func TestDeepseekJoke(t *testing.T) {
	apiKey, err := loadAPIKey()
	if err != nil || apiKey == "" {
		t.Skip("No API key configured, skipping live Deepseek integration test")
	}

	client, err := deepseek.NewClient(apiKey)
	if err != nil {
		t.Fatalf("Failed to create Deepseek client: %v", err)
	}

	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_CHAT_MODEL,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: "Tell me a random short joke.",
			},
		},
	}

	resp, err := client.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		t.Fatalf("Deepseek API call failed: %v", err)
	}

	if len(resp.Choices) == 0 {
		t.Fatalf("No choices returned from Deepseek API")
	}

	joke := resp.Choices[0].Message.Content
	if joke == "" {
		t.Fatalf("Empty joke returned from Deepseek API")
	}

	t.Logf("Received joke: %s", joke)
}
