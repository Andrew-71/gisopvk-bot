package service

import (
	"log/slog"

	// "github.com/bmstu-itstech/apollo/internal/app"
	// "github.com/bmstu-itstech/apollo/internal/app/command"
	// "github.com/bmstu-itstech/apollo/internal/app/query"
	// "github.com/bmstu-itstech/apollo/internal/common/logs"
	// "github.com/bmstu-itstech/apollo/internal/domain/material"
	"github.com/Andrew-71/gisopvk-bot/internal/app"
	"github.com/Andrew-71/gisopvk-bot/internal/app/query"
	"github.com/Andrew-71/gisopvk-bot/internal/common/logs"
	"github.com/Andrew-71/gisopvk-bot/internal/domain"
	mock_bot "github.com/Andrew-71/gisopvk-bot/internal/infra/mock"
)

type Cleanup func()

func NewApplication() (*app.Application, Cleanup) {
	logger := logs.DefaultLogger()

	// api := os.Getenv("BOT_URI")
	// bot := ai_bot.NewAiBot(api)
	bot := mock_bot.NewMockBot()
	return newApplication(logger, bot), func() {}
}

func NewTestApplication() (*app.Application, *mock_bot.MockBot) {
	logger := logs.DefaultLogger()
	bot := mock_bot.NewMockBot()
	return newApplication(logger, bot), &bot
}

func newApplication(
	logger *slog.Logger,
	bot domain.Bot,
) *app.Application {
	return &app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetReply: query.NewGetReplyHandler(bot, logger),
		},
	}
}
