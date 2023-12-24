package main

import (
	"context"
	"os"

	"github.com/aqyuki/hololive_news_crawler/internal/cli"
	"github.com/aqyuki/hololive_news_crawler/pkg/logging"
)

func main() {
	logger := logging.NewLogger()
	ctx := logging.WithLogger(context.Background(), logger)

	c := cli.NewCrawlCommand(ctx)
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
