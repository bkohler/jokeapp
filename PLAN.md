# Plan: Go Command-Line Joke App (`jokeapp`) - TDD Approach

This document outlines the plan for developing the `jokeapp` command-line tool using a Test-Driven Development (TDD) approach with live integration testing for the Deepseek API.

**I. Project Setup & Dependencies**

1.  **Project Directory:**
    ```bash
    mkdir jokeapp
    cd jokeapp
    ```
2.  **Go Module:**
    ```bash
    go mod init jokeapp
    ```
3.  **Dependencies:**
    ```bash
    go get github.com/cohesion-org/deepseek-go
    go get github.com/spf13/viper
    ```
    *Standard libraries (`flag`, `fmt`, `os`, `os/user`, `path/filepath`, `strings`, `bufio`, `context`, `testing`) will also be used.*

**II. Development Cycle (TDD)**

For each functional component listed below, the development cycle will follow these steps:
1.  **Write Test:** Create a test file (e.g., `component_test.go`) and write a failing test case for the desired functionality.
2.  **Write Code:** Implement the minimal code required in the corresponding source file (e.g., `main.go`, `config.go`) to make the test pass.
3.  **Refactor:** Improve the code structure and clarity while ensuring all tests continue to pass.

**III. Functional Components & Testing Strategy**

1.  **API Key Management (`~/.jokeapp.yaml`)**
    *   **Functionality:** Read API key from `~/.jokeapp.yaml`. If the file or key is missing, prompt the user to enter it and save it to the file.
    *   **Unit Tests (`config_test.go`):**
        *   Test config file path generation.
        *   Test reading from a mock config file (using temporary files/mock filesystem).
        *   Test the prompt-and-save logic (mocking `os.Stdin`, `os.Stdout`, filesystem/viper writes).
        *   Test handling of empty/invalid input during the prompt.
    *   **Implementation (`config.go` / `main.go`):** Use `os/user`, `path/filepath`, `viper`, `bufio`, `fmt`, `strings`.

2.  **Command-Line Flags**
    *   **Functionality:** Define and parse flags: `-h` (help), `-g` (geek), `-n` (nerd), `-s` (sport), `-r` (run). Handle `-h` to display usage.
    *   **Unit Tests (`flags_test.go`):**
        *   Test parsing various flag combinations.
        *   Test help message generation/output (capture stdout).
    *   **Implementation (`flags.go` / `main.go`):** Use the standard `flag` package.

3.  **Joke Category Determination**
    *   **Functionality:** Collect specified categories based on parsed flags. Default to "random" if no category flags are set.
    *   **Unit Tests (`categories_test.go`):**
        *   Test the default "random" case (no flags).
        *   Test single category flag scenarios.
        *   Test multiple category flag scenarios.
    *   **Implementation (`categories.go` / `main.go`):** Logic to check flag values and build a category list/string.

4.  **Deepseek Prompt Construction**
    *   **Functionality:** Create the appropriate prompt string for the Deepseek API based on the determined categories (random, single, or combined).
    *   **Unit Tests (`prompt_test.go`):**
        *   Test prompt generation for the "random" case.
        *   Test prompt generation for single categories.
        *   Test prompt generation for combined categories.
    *   **Implementation (`prompt.go` / `main.go`):** Use `fmt.Sprintf` and `strings.Join`.

5.  **Deepseek API Interaction (Integration Test)**
    *   **Functionality:** Initialize the Deepseek client, send a request, receive the response, and parse the joke.
    *   **Integration Tests (`deepseek_integration_test.go`):**
        *   **Tag:** Use `//go:build integration` build tag.
        *   **Prerequisite:** Requires a valid API key in `~/.jokeapp.yaml`. Tests should skip (`t.Skip()`) if the key is unavailable.
        *   **Test:**
            *   Initialize the *real* `deepseek.Client`.
            *   Send a request (e.g., for a random joke).
            *   Assert `err` is `nil`.
            *   Assert response structure is valid (e.g., `resp.Choices` is not empty).
            *   Assert joke content (`resp.Choices[0].Message.Content`) is not empty.
            *   *(Do not assert specific joke content)*.
    *   **Implementation (`deepseek_client.go` / `main.go`):** Use `github.com/cohesion-org/deepseek-go` and `context`.

6.  **Main Application Logic (`main.go`)**
    *   **Functionality:** Orchestrate all steps: get API key, parse flags, handle help, determine categories, build prompt, call API, print result or error.
    *   **Testing:** Primarily covered by component and integration tests. End-to-end tests (running the compiled binary) can be added later if needed.

**IV. Running Tests**

*   **Run Unit Tests:**
    ```bash
    go test ./...
    ```
*   **Run Integration Tests (Requires API Key):**
    ```bash
    go test -tags=integration ./...
    ```
*   **Run All Tests:**
    ```bash
    go test -tags=integration ./...
    ```

**V. Workflow Diagram (Conceptual TDD Flow)**

```mermaid
graph TD
    subgraph "Component: Config"
        T1[Write Config Test] --> C1[Write Config Code] --> R1[Refactor Config]
    end
    subgraph "Component: Flags"
        T2[Write Flags Test] --> C2[Write Flags Code] --> R2[Refactor Flags]
    end
    subgraph "Component: Categories"
        T3[Write Categories Test] --> C3[Write Categories Code] --> R3[Refactor Categories]
    end
    subgraph "Component: Prompt"
        T4[Write Prompt Test] --> C4[Write Prompt Code] --> R4[Refactor Prompt]
    end
    subgraph "Component: Deepseek API"
        T5[Write Integration Test] --> C5[Write API Code] --> R5[Refactor API Code]
    end
    subgraph "Main Orchestration"
        T6[Assemble in main_test.go (Optional E2E)] --> C6[Write main.go Code] --> R6[Refactor main.go]
    end

    R1 --> T2
    R2 --> T3
    R3 --> T4
    R4 --> T5
    R5 --> C6