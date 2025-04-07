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

	// Additional joke category flags
	dadFlag := flag.Bool("d", false, "Dad joke")
	animalFlag := flag.Bool("a", false, "Animal joke")
	techFlag := flag.Bool("t", false, "Tech joke")
	programmerFlag := flag.Bool("p", false, "Programmer joke")
	mathFlag := flag.Bool("m", false, "Math joke")
	scienceFlag := flag.Bool("c", false, "Science joke")
	foodFlag := flag.Bool("f", false, "Food joke")
	doctorFlag := flag.Bool("o", false, "Doctor joke")
	lawyerFlag := flag.Bool("l", false, "Lawyer joke")
	politicalFlag := flag.Bool("q", false, "Political joke")
	blondeFlag := flag.Bool("b", false, "Blonde joke")
	knockFlag := flag.Bool("k", false, "Knock-knock joke")
	schoolFlag := flag.Bool("s2", false, "School joke")
	musicFlag := flag.Bool("u", false, "Music joke")
	movieFlag := flag.Bool("v", false, "Movie joke")
	historyFlag := flag.Bool("h2", false, "History joke")
	relationshipFlag := flag.Bool("r2", false, "Relationship joke")
	workFlag := flag.Bool("w", false, "Work joke")
	travelFlag := flag.Bool("j", false, "Travel joke")
	punFlag := flag.Bool("n2", false, "Pun joke")

	flag.Parse()

	if *helpFlag {
		fmt.Println("Usage: jokeapp [options]")
		fmt.Println("Options:")
		fmt.Println("  -h        Show help")
		fmt.Println("  -g        Geek joke")
		fmt.Println("  -n        Nerd joke")
		fmt.Println("  -s        Sport joke")
		fmt.Println("  -r        Running joke")
		fmt.Println("  -d        Dad joke")
		fmt.Println("  -a        Animal joke")
		fmt.Println("  -t        Tech joke")
		fmt.Println("  -p        Programmer joke")
		fmt.Println("  -m        Math joke")
		fmt.Println("  -c        Science joke")
		fmt.Println("  -f        Food joke")
		fmt.Println("  -o        Doctor joke")
		fmt.Println("  -l        Lawyer joke")
		fmt.Println("  -q        Political joke")
		fmt.Println("  -b        Blonde joke")
		fmt.Println("  -k        Knock-knock joke")
		fmt.Println("  -s2       School joke")
		fmt.Println("  -u        Music joke")
		fmt.Println("  -v        Movie joke")
		fmt.Println("  -h2       History joke")
		fmt.Println("  -r2       Relationship joke")
		fmt.Println("  -w        Work joke")
		fmt.Println("  -j        Travel joke")
		fmt.Println("  -n2       Pun joke")
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
	if *dadFlag {
		categories = append(categories, "dad")
	}
	if *animalFlag {
		categories = append(categories, "animal")
	}
	if *techFlag {
		categories = append(categories, "tech")
	}
	if *programmerFlag {
		categories = append(categories, "programmer")
	}
	if *mathFlag {
		categories = append(categories, "math")
	}
	if *scienceFlag {
		categories = append(categories, "science")
	}
	if *foodFlag {
		categories = append(categories, "food")
	}
	if *doctorFlag {
		categories = append(categories, "doctor")
	}
	if *lawyerFlag {
		categories = append(categories, "lawyer")
	}
	if *politicalFlag {
		categories = append(categories, "political")
	}
	if *blondeFlag {
		categories = append(categories, "blonde")
	}
	if *knockFlag {
		categories = append(categories, "knock-knock")
	}
	if *schoolFlag {
		categories = append(categories, "school")
	}
	if *musicFlag {
		categories = append(categories, "music")
	}
	if *movieFlag {
		categories = append(categories, "movie")
	}
	if *historyFlag {
		categories = append(categories, "history")
	}
	if *relationshipFlag {
		categories = append(categories, "relationship")
	}
	if *workFlag {
		categories = append(categories, "work")
	}
	if *travelFlag {
		categories = append(categories, "travel")
	}
	if *punFlag {
		categories = append(categories, "pun")
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
