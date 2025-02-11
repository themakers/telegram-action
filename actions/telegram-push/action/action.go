package action

import (
	"context"
	"os"
	"strings"
)

type Options struct {
	BotToken string
	ChatID   string
}

// https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/store-information-in-variables#default-environment-variables

func Invoke(ctx context.Context, opt Options) {

	notify(ctx, notifyOptions{
		BotToken:        opt.BotToken,
		ChatID:          opt.ChatID,
		MessageTemplate: strings.Trim(os.Getenv("MESSAGE_TEMPLATE"), " \n\r\t"),

		Repository: os.Getenv("GITHUB_REPOSITORY"),
		Actor:      os.Getenv("GITHUB_ACTOR"),
		Branch:     os.Getenv("GITHUB_REF_NAME"),
		CommitSHA:  os.Getenv("GITHUB_SHA"),
	})

}
