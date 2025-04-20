package service

import (
	"log/slog"
	"os"

	"github.com/Andrew-71/gisopvk-bot/internal/app"
	"github.com/Andrew-71/gisopvk-bot/internal/app/query"
	"github.com/Andrew-71/gisopvk-bot/internal/common/decorator"
	"github.com/Andrew-71/gisopvk-bot/internal/common/logs"
	"github.com/Andrew-71/gisopvk-bot/internal/common/metrics"
	"github.com/Andrew-71/gisopvk-bot/internal/domain"
	"github.com/Andrew-71/gisopvk-bot/internal/infra/ai"
	mock_bot "github.com/Andrew-71/gisopvk-bot/internal/infra/mock"
)

type Cleanup func()

func NewApplication() (*app.Application, Cleanup) {
	// FIXME: Temporarily needs hardcoding :(
	// bot := ai.NewAiBot("http://localhost:11434/api/chat", "llama3.2")
	bot := ai.NewAiBot(os.Getenv("AI_URL"), os.Getenv("AI_MODEL"))
	logger := logs.DefaultLogger()
	client := metrics.NewPromMetrics()
	return newApplication(bot, logger, client), func() {}
}

func NewTestApplication() (*app.Application, *mock_bot.MockBot) {
	bot := mock_bot.NewMockBot()
	logger := logs.DefaultLogger()
	client := metrics.NewPromMetrics()
	return newApplication(bot, logger, client), &bot
}

func newApplication(
	bot domain.Bot,
	logger *slog.Logger,
	metricsClient decorator.MetricsClient,
) *app.Application {
	return &app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetReply: query.NewGetReplyHandler(bot, logger, metricsClient),
		},
	}
}
