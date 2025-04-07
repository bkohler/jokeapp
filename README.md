# jokeapp

A Go command-line tool that fetches AI-generated jokes from Deepseek, supporting multiple joke categories and built with Test-Driven Development.

## Features

- Fetches jokes from Deepseek AI
- Supports joke categories:
  - Geek (`-g`)
  - Nerd (`-n`)
  - Sport (`-s`)
  - Running (`-r`)
- Combine multiple categories for multi-themed jokes
- Prompts for Deepseek API key on first run and saves it securely
- Includes unit and integration tests
- GitHub Actions CI for automated testing

## Usage

Build the app:

```bash
./build.sh
```

Run the app:

```bash
./jokeapp [flags]
```

Flags:

- `-h` Show help
- `-g` Geek joke
- `-n` Nerd joke
- `-s` Sport joke
- `-r` Running joke

Examples:

```bash
./jokeapp -g
./jokeapp -n -s
./jokeapp -g -n -s -r
```

## API Key Setup

On first run, you will be prompted to enter your Deepseek API key. It will be saved in `~/.jokeapp.yaml`.

## Testing

Run all tests:

```bash
go test ./...
```

Run integration tests (requires API key):

```bash
go test -tags=integration ./...
```

## GitHub Actions

The project includes a GitHub Actions workflow that:

- Runs unit tests on every push and pull request
- Runs integration tests if the `DEEPSEEK_API_KEY` secret is set

## License

MIT