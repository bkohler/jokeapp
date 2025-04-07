package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
	"github.com/spf13/viper"
)

func main() {
	// Load or prompt for API key
	apiKey, err := loadAPIKey()
	if err != nil || apiKey == "" {
		fmt.Println("Deepseek API key not found. Please enter your API key:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(input)
		if apiKey == "" {
			fmt.Fprintln(os.Stderr, "API key cannot be empty. Exiting.")
			os.Exit(1)
		}
		viper.Set("deepseek_api_key", apiKey)
		configPath := getConfigFilePath()
		err := viper.WriteConfigAs(configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to save API key: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("API key saved successfully.")
	}

	// Parse flags
	helpFlag := flag.Bool("h", false, "Display help")
	geekFlag := flag.Bool("g", false, "Geek joke")
	nerdFlag := flag.Bool("n", false, "Nerd joke")
	sportFlag := flag.Bool("s", false, "Sport joke")
	runFlag := flag.Bool("r", false, "Running joke")
	flag.Parse()

	if *helpFlag {
		fmt.Println("Usage: jokeapp [options]")
		fmt.Println("Options:")
		fmt.Println("  -h        Show help")
		fmt.Println("  -g        Geek joke")
		fmt.Println("  -n        Nerd joke")
		fmt.Println("  -s        Sport joke")
		fmt.Println("  -r        Running joke")
		os.Exit(0)
	}

	// Determine categories
	categories := []string{}
	if *geekFlag {
		categories = append(categories, "geek")
	}
	if *nerdFlag {
		categories = append(categories, "nerd")
	}
	if *sportFlag {
		categories = append(categories, "sport")
	}
	if *runFlag {
		categories = append(categories, "running")
	}

	// Build prompt
	prompt := buildPrompt(categories)

	// Initialize Deepseek client
	client, err := deepseek.NewClient(apiKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Deepseek client: %v\n", err)
		os.Exit(1)
	}

	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_CHAT_MODEL,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	resp, err := client.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Deepseek API call failed: %v\n", err)
		os.Exit(1)
	}

	if len(resp.Choices) == 0 {
		fmt.Fprintln(os.Stderr, "No joke received from Deepseek API.")
		os.Exit(1)
	}

	joke := resp.Choices[0].Message.Content
	fmt.Println(joke)
}
