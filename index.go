package main

import (
	"fmt"
	"log"

	/*"os/exec"
	"strings"*/

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    bot, err := tgbotapi.NewBotAPI("Enter your token here")
    if err != nil {
        log.Fatal(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

		if update.Message.IsCommand() && update.Message.Command() == "start" {
            name := "Vlas Kozlov" // Здесь можно присвоить значение переменной
			class := "anus"
            msgText := fmt.Sprintf("Привет, %s! Добро пожаловать! %s", name, class)
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
            bot.Send(msg)

        /*msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
        bot.Send(msg)*/
    

	/*command := strings.TrimSpace(update.Message.Text)

        if strings.HasPrefix(command, "/eval") {
            code := strings.TrimSpace(strings.TrimPrefix(command, "/eval"))

            output, err := exec.Command("go", "run", "-e", code).CombinedOutput()
            if err != nil {
                log.Println(err)
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка выполнения кода.")
                bot.Send(msg)
            } else {
                result := string(output)
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
                bot.Send(msg)
            }*/
        }
	}
}
