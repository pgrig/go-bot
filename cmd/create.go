/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	
	tele "gopkg.in/telebot.v3"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command creates a telegram bot",
	Long: `This command creates and runs a Telegram bot.
You can provide the bot token either through environment variable
TELEGRAM_BOT_TOKEN or through the --token flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Telegram bot...")
		
		// First, check environment variable
		token := os.Getenv("TELEGRAM_BOT_TOKEN")
		
		// If environment variable is not set, use the flag
		if token == "" {
			tokenFlag, _ := cmd.Flags().GetString("token")
			if tokenFlag != "" {
				token = tokenFlag
			}
		}
		
		// Verify that we have a token
		if token == "" {
			fmt.Println("Error: Bot token is required. Set TELEGRAM_BOT_TOKEN environment variable or use --token flag.")
			return
		}
		
		// Bot settings
		pref := tele.Settings{
			Token:  token,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}
		
		// Create the bot
		b, err := tele.NewBot(pref)
		if err != nil {
			fmt.Printf("Error creating bot: %v\n", err)
			return
		}
		
		// Handler for /start command
		b.Handle("/start", func(c tele.Context) error {
			return c.Send("Hello! I'm your new bot. Use /help to get a list of commands.")
		})
		
		// Handler for /help command
		b.Handle("/help", func(c tele.Context) error {
			return c.Send(`Available commands:
/start - Start the bot
/help - Show this message
You can also send me any message, and I'll respond`)
		})
		
		// Handler for regular text messages
		b.Handle(tele.OnText, func(c tele.Context) error {
			msg := c.Message()
			fmt.Printf("Received message from %s: %s\n", msg.Sender.Username, msg.Text)
			return c.Send("You wrote: " + msg.Text)
		})
		
		fmt.Println("Bot is running. Press Ctrl+C to stop.")
		
		// Start the bot
		b.Start()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Add flag for bot token, but don't make it required
	createCmd.Flags().StringP("token", "t", "", "Telegram Bot Token (can also be set via TELEGRAM_BOT_TOKEN environment variable)")
}