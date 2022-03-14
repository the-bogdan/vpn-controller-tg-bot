package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	updatesOffset  = 0
	updatesLimit   = 0
	updatesTimeout = 60
)

type Bot struct {
	api          *tgbotapi.BotAPI
	allowedUsers map[int64]bool
}

func NewBot(token string, allowedUsers map[int64]bool) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{
		api:          bot,
		allowedUsers: allowedUsers,
	}, nil
}

func (b Bot) Run() {
	updatesChan := b.api.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  updatesOffset,
		Limit:   updatesLimit,
		Timeout: updatesTimeout,
	})

	for update := range updatesChan {
		if isValid := b.validate(update); !isValid {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID
		b.api.Send(msg)
	}
}

func (b Bot) validate(update tgbotapi.Update) bool {
	// Проверяем что обновление это сообщение
	if update.Message == nil {
		return false
	}
	// Проверяем что у пользователя есть права пользоваться ботом
	if _, ok := b.allowedUsers[update.Message.From.ID]; !ok {
		msg := "У вас нет прав пользоваться этим ботом"
		b.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))

		log.Printf(
			"user [%d || @%s || %s %s] has no access to use bot",
			update.Message.From.ID,
			update.Message.From.UserName,
			update.Message.From.FirstName,
			update.Message.From.LastName,
		)
		return false
	}
	return true
}
