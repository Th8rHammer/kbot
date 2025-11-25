package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v4"
)

// ----------------------------------------------------
// A. Налаштування Cobra
// ----------------------------------------------------

// Ініціалізація кореневої команди, яка буде запускати бота
var rootCmd = &cobra.Command{
	Use:   "kbot",
	Short: "A functional Telegram bot built with Go.",
	Long: `A Telegram bot that handles messages and provides basic commands.
Built using Go, telebot.v4, and cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		startBot() // Запуск логіки бота
	},
}

// Функція для виконання кореневої команди, повертає помилку для обробки в main.go
func Execute() error {
	return rootCmd.Execute()
}

// ----------------------------------------------------
// B. Логіка Ініціалізації та Запуску Бота (startBot)
// ----------------------------------------------------

func startBot() {
	// 1. Отримання токена
	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		// Вихід, якщо токен не встановлено
		log.Fatal("TELE_TOKEN environment variable not set. Please set your bot token.")
	}

	// 2. Налаштування бота
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second}, 
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
		return
	}

	log.Printf("Bot initialized successfully! Running as @%s", b.Me.Username)

	// 3. Реєстрація обробників
	b.Handle("/start", handleStart(b))
    b.Handle("/settings", handleSettings(b))
	b.Handle(tele.OnText, handleText(b)) 
    
	// 4. Запуск бота
	b.Start()
}

// ----------------------------------------------------
// C. Функції-Обробники (Handlers)
// ----------------------------------------------------

// handleStart обробляє команду /start
func handleStart(b *tele.Bot) tele.HandlerFunc {
	return func(c tele.Context) error {
		welcomeMessage := fmt.Sprintf("👋 Привіт, %s! Я kbot. Надішли мені будь-яке повідомлення.", c.Sender().FirstName)
		return c.Send(welcomeMessage) 
	}
}

// handleSettings обробляє команду /settings
func handleSettings(b *tele.Bot) tele.HandlerFunc {
	return func(c tele.Context) error {
		settingsMessage := "⚙️ Тут будуть налаштування вашого бота. Доступні команди: /start, /settings."
		return c.Send(settingsMessage)
	}
}

// handleText обробляє будь-яке текстове повідомлення
func handleText(b *tele.Bot) tele.HandlerFunc {
	return func(c tele.Context) error {
		userText := c.Text()
		
		var response string
		
		// Приклад логіки відповіді
		if userText == "команда" {
			response = "Викликано спеціальну команду! Дякую."
		} else {
			// Відповідь-ехо
			response = fmt.Sprintf("Ти написав: \"%s\"", userText)
		}
		
		return c.Reply(response)
	}
}