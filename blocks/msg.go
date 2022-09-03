package blocks

import (
	"fmt"

	"github.com/Nikit-S/tgfr/template"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Msg struct {
	Text string
}

// Sends message to user with text m.Text
func (m Msg) Execute(bot *template.Bot, user *template.User) {
	msg := tgbotapi.NewMessage(user.GetUserId(), m.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.GetApi().Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}

type RepeatInput struct{}

// Sends message to user with last text input from user
func (m RepeatInput) Execute(bot *template.Bot, user *template.User) {
	update := <-user.GetChan()
	msg := tgbotapi.NewMessage(user.GetUserId(), update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.GetApi().Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}
