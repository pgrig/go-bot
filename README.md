# Telegram Bot CLI

A command-line interface tool for creating and managing Telegram bots built with Go and Cobra.

## Overview

This project provides a simple CLI to create, configure, and run Telegram bots. It's designed to be easy to use while following best practices for security and configuration.

## Installation

### Prerequisites

- Go 1.16 or later
- A Telegram bot token (get one from [@BotFather](https://t.me/BotFather))

### Building from source

```bash
# Clone the repository
git clone https://github.com/yourusername/telegram-bot-cli.git
cd telegram-bot-cli

# Build the project
go build -o telegram-bot-cli

# Make it globally available (optional)
sudo mv telegram-bot-cli /usr/local/bin/
```

## Usage

### Creating and starting a bot

```bash
# Using environment variable (recommended)
export TELEGRAM_BOT_TOKEN="your_token_here"
telegram-bot-cli create

# OR using command-line flag
telegram-bot-cli create --token="your_token_here"
telegram-bot-cli create -t "your_token_here"
```

## Features

- Simple command-line interface for bot management
- Secure token handling via environment variables
- Built-in command handlers for basic functionality
- Real-time message logging

## Configuration

### Bot Token

The bot token can be provided in two ways:

1. **Environment Variable** (recommended for security):
   ```bash
   export TELEGRAM_BOT_TOKEN="your_token_here"
   ```

2. **Command-line Flag**:
   ```bash
   --token="your_token_here" or -t "your_token_here"
   ```

Priority is given to the environment variable if both are set.

## Command Reference

### create

Creates and starts a new Telegram bot instance:

```bash
telegram-bot-cli create
```

#### Options:
- `--token, -t`: Specify the Telegram bot token

## Default Bot Commands

The bot comes with the following pre-configured commands:

- `/start` - Initiates the bot conversation with a welcome message
- `/help` - Shows available commands and usage instructions

## Development

### Project Structure

```
telegram-bot-cli/
├── cmd/             # Command definitions
│   ├── root.go      # Root command
│   ├── create.go    # Create command
│   └── ...
├── main.go          # Entry point
└── ...
```

### Adding New Commands

To add a new command to the bot:

1. Create a new file in the `cmd` directory
2. Use Cobra's command structure
3. Add the command to the root command in the `init()` function

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [telebot.v3](https://github.com/tucnak/telebot) - Telegram bot framework

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.