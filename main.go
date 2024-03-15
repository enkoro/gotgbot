package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("7126667493:AAFREt6WrI8CUiogTyBlen1-klLdEWqLCL8", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, bot *bot.Bot, update *models.Update) {
	fmt.Println(update.Message.Text)
	fmt.Println(update.Message.Chat.ID)
}
